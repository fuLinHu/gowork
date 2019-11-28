package main

import "fmt"

func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		fmt.Println("向通道接发送的数据为",i)
		out <- i
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}
func printer(in <-chan int) {
	for i := range in {
		fmt.Println("从通道接受到的数据为",i)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go printer(ch1)
	go squarer(ch2, ch1)
	printer(ch2)
}