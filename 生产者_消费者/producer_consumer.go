package main

import (
	"fmt"
	"time"
)

/*
__author__ = 'lawtech'
__date__ = '2018/8/22 上午11:31'
*/

func producer(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("Producer send:", i)
	}
}

func consumer(ch chan int) {
	for i := 0; i < 10; i++ {
		v := <-ch
		fmt.Println("Consumer received:", v)
	}
}

func main() {
	ch := make(chan int, 10)
	go producer(ch)
	go consumer(ch)
	time.Sleep(1 * time.Second)
}
