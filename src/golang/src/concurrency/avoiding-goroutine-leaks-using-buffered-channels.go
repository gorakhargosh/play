package main

import (
	"fmt"
	"net"
	"time"
)

func sendMsg(msg, addr string) error {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	defer conn.Close()
	_, err = fmt.Fprint(conn, msg)
	return err
}

// SHOW1 OMIT
func broadcastMsg(msg string, addrs []string) error { // HL
	errc := make(chan error, len(addrs)) // HL
	for _, addr := range addrs {
		go func(addr string) {
			errc <- sendMsg(msg, addr) // non-blocking, non-leaky. // HL
			fmt.Println("done")
		}(addr)
	}
	for _ = range addrs {
		if err := <-errc; err != nil {
			return err
		}
	}
	return nil
}

func main() { // HL
	addrs := []string{"localhost:8080", "google.com:80"}
	err := broadcastMsg("hi", addrs)
	time.Sleep(time.Second)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("everything went fine")
}

// END SHOW1 OMIT
