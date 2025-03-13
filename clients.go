package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/panjf2000/ants/v2"
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
	wg1.Wait()
}

func antsClients() {
	wg1.Add(concurrencyClients)
	pool, _ := ants.NewPool(50)
	defer pool.Release()
	for i := 1; i <= concurrencyClients; i++ {
		r := &Request{
			args:       []int{i},
			resultChan: make(chan int),
		}
		_ = pool.Submit(func() {
			ClientReq(r)
		})
	}
	wg1.Wait()
}

func ClientReq(r *Request) {
	defer wg1.Done()
	var start = time.Now()
	queue <- r
	//	go func() {
	res := <-r.resultChan
	fmt.Printf("current args is %d, the result is %d, 耗时： %d ms \n", r.args[0], res, time.Since(start).Milliseconds())
	// }()
}
