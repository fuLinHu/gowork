package main

import (
	"fmt"
)

// 生产者
func produce2(ch chan<- int) {
	for i := 0; i < 1; i++ {
		ch <- i
		fmt.Println("生产了", i)
	}
}

// 消费者
func custom2(ch <-chan int, ch2 chan<- int) {
	for i := range ch {
		fmt.Println("消费了", i)
		ch2 <- i
	}

}

func main() {
	ch := make(chan int)
	ch2 := make(chan int, 100)
	go custom2(ch, ch2)
	produce2(ch)
	<-ch2
}
