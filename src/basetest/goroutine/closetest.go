package main

import "fmt"

func main() {
	ch := make(chan int, 10)
	ch <- 1
	close(ch)
	x := <-ch
	fmt.Println(x)
}
