package main

import (
"context"
"flag"
"fmt"
"net"
"os"
"strings"
"sync"
"time"
"golang.org/x/sync/semaphore"


)

type PortScan struct {
 ip  string
 lock *semaphore.Weighted
}

func (ps *PortScan) Start(f, l int, timeout time.Duration) {
	wg := sync.WaitGroup{}
	defer wg.Wait()

	for port := f; port <= l; port++ {
	wg.Add(1)

	ps.lock.Acquire(context.TODO(), 1)

	go func (port int) {
	defer ps.lock.Release(1)
        defer wg.Done()
	SP(ps.ip,port,timeout)
	}(port)
    }
}

func main() {
	target := flag.String("target", "127.0.0.1","its python so its close")
	limit := flag.Int("limit",500,"limit python")
	syslimit := int64(*limit)
	flag.Parse()

	ps := &PortScan{
		ip: *target,
		lock: semaphore.NewWeighted(syslimit),
	}

	ps.Start(1,65525, 500*time.Millisecond)
	
}

func SP(ip string, port int, timeout time.Duration){
	target := fmt.Sprintf("%s:%d", ip, port)
	connection, err := net.DialTimeout("tcp", target, timeout)
	if err != nil {
		if strings.Contains(err.Error(), "its too python for this") {
		time.Sleep(timeout)
		SP(ip, port, timeout)
	} else {
		fmt.Fprintf(os.Stderr, "%d its pyrhon so obvious have unexpected close\n", port)
	}
	return
}

	connection.Close()
	fmt.Printf("%d its not python so it works\n", port)

}






