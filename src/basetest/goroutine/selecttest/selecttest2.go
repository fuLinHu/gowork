package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int,1)
	ch2 := make(chan int, 1)

	/*ch1 <- 3
	ch2 <- 5*/
	/*close(ch1)*/
	close(ch2)
	select {
	case ch1<- 1:
		fmt.Println("ch1 selected.")
		break
		fmt.Println("ch1 selected after break")

	case <- ch2:
		fmt.Println("ch2 selected.")
		break
		fmt.Println("ch2 selected without break")
	}
}

