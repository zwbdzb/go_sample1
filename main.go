package main

import (
	"fmt"
	"time"
)

var concurrencyClients = 1000
var queueLength = 100
var queue = make(chan *Request, queueLength) // 请求队列长度
var Maxoutstanding int = 10                  // 服务器并发受限10

func main() {
	go server(queue)
	var start = time.Now()

	// clients()
	antsClients()
	fmt.Printf("客户端并发%d请求，服务器请求队列长度%d，服务器限流%d，总共耗时%d ms \n", concurrencyClients, queueLength, Maxoutstanding, time.Since(start).Milliseconds())
}
