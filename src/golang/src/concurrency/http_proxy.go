// From: Building Services in Go - Mark Smith.
// Watch: https://www.youtube.com/watch?v=MeOK1UzGHYw
package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const hostAddr = ":8080"
const beHostAddr = "127.0.0.1:8081"

type Backend struct {
	net.Conn
	Reader *bufio.Reader
	Writer *bufio.Writer
}

type Stats struct {
	mu            sync.Mutex
	responseBytes map[string]int64
}

func (s *Stats) update(req *http.Request, resp *http.Response) int64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	bytes := s.responseBytes[req.URL.Path] + resp.ContentLength
	s.responseBytes[req.URL.Path] = bytes
	return bytes
}

var stats Stats
var beq chan *Backend

func init() {
	stats = Stats{responseBytes: make(map[string]int64)}
	beq = make(chan *Backend, 10)
}

func queueBackend(q chan *Backend, be *Backend) {
	select {
	case q <- be:
		// be reenqueued.
	case <-time.After(1 * time.Second):
		be.Close()
	}
}

func getBackend(q chan *Backend, beHostAddr string) (*Backend, error) {
	select {
	case be := <-q:
		return be, nil
	case <-time.After(100 * time.Millisecond):
		be, err := net.Dial("tcp", beHostAddr)
		if err != nil {
			return nil, err
		}
		return &Backend{
			Conn:   be,
			Reader: bufio.NewReader(be),
			Writer: bufio.NewWriter(be),
		}, nil
	}
}

// proxy1 is a serial implementation of an HTTP proxy.
func proxy1() {
	if listener, err := net.Listen("tcp", hostAddr); err == nil {
		for {
			if client, err := listener.Accept(); err == nil { // Blocking.
				reader := bufio.NewReader(client)
				if req, err := http.ReadRequest(reader); err == nil {
					if be, err := net.Dial("tcp", beHostAddr); err == nil {
						backendReader := bufio.NewReader(be)
						if err := req.Write(be); err == nil {
							if resp, err := http.ReadResponse(backendReader, req); err == nil {
								resp.Close = true
								if err := resp.Write(client); err == nil {
									log.Printf("%s: %d", req.URL.Path, resp.StatusCode)
								}
								client.Close()
								// Accept next connection.
							}
						}
					}
				}
			}
		}
	}
}

// handleClient handles a connection given to it by a listener.Accept call.
func handleClient(client net.Conn) {
	defer client.Close()
	reader := bufio.NewReader(client)

	for {
		req, err := http.ReadRequest(reader)
		if err != nil {
			if err != io.EOF {
				log.Printf("Failed to read request: %s", err)
			}
			return
		}
		if be, err := net.Dial("tcp", beHostAddr); err == nil {
			backendReader := bufio.NewReader(be)
			if err := req.Write(be); err == nil {
				if resp, err := http.ReadResponse(backendReader, req); err == nil {
					bytes := stats.update(req, resp)
					resp.Header.Set("X-Bytes", strconv.FormatInt(bytes, 10))

					if err := resp.Write(client); err == nil {
						log.Printf("%s: %d", req.URL.Path, resp.StatusCode)
					}
				}
			}
		}
	}
}

// proxy2 is a concurrent implementation of an HTTP proxy.
func proxy2() {
	listener, err := net.Listen("tcp", hostAddr)
	if err != nil {
		log.Fatalf("Failed to listen %q -- error: %q", hostAddr, err)
	}
	for {
		if client, err := listener.Accept(); err == nil {
			go handleClient(client)
		}
	}
}

// handleClientPooled handles a connection given to it by a listener.Accept call
// and uses a pool of connections to the backend to handle requests.
func handleClientPooled(client net.Conn) {
	defer client.Close()
	reader := bufio.NewReader(client)

	for {
		req, err := http.ReadRequest(reader)
		if err != nil {
			if err != io.EOF {
				log.Printf("Failed to read request: %s", err)
			}
			return
		}

		be, err := getBackend(beq, beHostAddr)
		if err != nil {
			return
		}

		if err := req.Write(be.Writer); err == nil {
			be.Writer.Flush()
			if resp, err := http.ReadResponse(be.Reader, req); err == nil {
				if err := resp.Write(client); err == nil {
					log.Printf("proxied %s: %d", req.URL.Path, resp.StatusCode)
				}
				if resp.Close {
					return
				}
			}
		}

		// The following may block for a second, so we don't make
		// handleConnectionQueue wait that long.
		go queueBackend(beq, be)
	}
}

// proxy3 is a concurrent implementation of an HTTP proxy.
func proxy3() {
	listener, err := net.Listen("tcp", hostAddr)
	if err != nil {
		log.Fatalf("Failed to listen %q -- error: %q", hostAddr, err)
	}
	for {
		if client, err := listener.Accept(); err == nil {
			go handleClientPooled(client)
		}
	}
}

func main() {
	proxy3()
}
