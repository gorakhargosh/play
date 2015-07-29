package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"sync"
)

type Task interface {
	Process()
	Print()
}

type TaskGenerator func(line string) Task

func run(r io.Reader, f TaskGenerator, workerCount int) {
	var wg sync.WaitGroup

	in := make(chan Task)
	wg.Add(1)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			if len(strings.TrimSpace(s.Text())) > 0 {
				in <- f(s.Text())
			}
		}
		if s.Err() != nil {
			log.Fatalf("Error reading input: %s", s.Err())
		}
		close(in)
		wg.Done()
	}()

	// Start writing the work output to the output channel.
	out := make(chan Task)
	wg.Add(workerCount)
	for i := 0; i < workerCount; i++ {
		go func() {
			for t := range in {
				t.Process()
				out <- t
			}
			wg.Done()
		}()
	}

	// Close the channel when all the workers have done their work.
	go func() {
		wg.Wait()
		close(out)
	}()

	for t := range out {
		t.Print()
	}
}

type lookupTask struct {
	query string
	ns    []*net.NS
	err   error
}

func (t *lookupTask) Process() {
	ns, err := net.LookupNS(t.query)
	if err != nil {
		log.Fatal(err)
		return
	}
	t.ns = ns
}

func (t *lookupTask) Print() {
	fmt.Printf("%s DNS servers: \n", t.query)
	for _, ns := range t.ns {
		fmt.Printf("%s\n", ns.Host)
	}
	fmt.Println()
}

func NewLookupTask(line string) Task {
	return &lookupTask{query: line}
}

func main() {
	run(os.Stdin, NewLookupTask, 100)
}
