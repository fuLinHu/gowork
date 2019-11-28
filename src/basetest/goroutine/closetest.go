package main

import "fmt"

func main() {
	ch := make(chan int)
	/*ch <- 1*/
	close(ch)
	x := <-ch
	fmt.Println(x)
}
