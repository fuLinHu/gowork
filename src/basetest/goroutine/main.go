package main

import (
	"fmt"
)

// 生产者
func produce1(ch chan<- int) {
	for i := 0; i < 1; i++ {
		ch <- i
		fmt.Println("生产了", i)
	}
}

// 消费者
func custom1(ch <-chan int) {
	for i := range ch {
		fmt.Println("消费了", i)
	}
}

func main() {
	ch := make(chan int)
	go custom1(ch)
	produce1(ch)
}
