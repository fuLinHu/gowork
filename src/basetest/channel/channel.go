package main

import "fmt"

func res(ch chan int) {
	x := <-ch
	fmt.Print(x)
}

func main() {
	var ch chan int
	fmt.Println(ch)

	ch = make(chan int, 1)
	fmt.Println(ch)

	ch <- 10
	//go res(ch)

	close(ch)
}
