package main

import (
	"fmt"
	"sync"
)

type Request struct {
	args       []int
	resultChan chan int
}

var wg1 sync.WaitGroup // 确保所有的客户端请求都已经发出且收到响应

func clients() {
	fmt.Printf("start %d concurrency client request\n ", concurrencyClients)
	for i := 1; i <= concurrencyClients; i++ {
		r := &Request{
			args:       []int{i},
			resultChan: make(chan int),
		}
		wg1.Add(1)
		go ClientReq(r)
	}
	wg1.Wait() // 到目前为止， golang运行时会检测到：剩余的50个协程都处于阻塞状态，并且没有其他协程可以解除，故会报死锁。

}
func ClientReq(r *Request) {
	defer wg1.Done()
	queue <- r
	go func() {
		res := <-r.resultChan
		fmt.Printf("current args is %d, the result is %d \n", r.args[0], res)
	}()
}
