/* 实现一个有请求队列功能的并发请求受限服务器*/

package main

import (
	"fmt"
	"sync"
	"time"
)

var sem = make(chan int, Maxoutstanding)

var wg2 sync.WaitGroup

func server(queue chan *Request) {
	fmt.Printf("Server is already, listen req \n")

	for req := range queue {
		req := req
		sem <- 1

		wg2.Add(1)
		go func() {
			defer wg2.Done()
			process(req)
			<-sem
		}()
	}
}

func process(req *Request) {
	s := sum(req.args)
	req.resultChan <- s
	//fmt.Printf("current args is %d, the result is %d \n", req.args[0], s)
}
func sum(a []int) (s int) {
	for i := 1; i <= a[0]; i++ {
		s += i
	}
	time.Sleep(time.Millisecond * 20)
	return s
}
