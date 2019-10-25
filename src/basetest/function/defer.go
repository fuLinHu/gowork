package main

import "fmt"

func main() {
	defer fmt.Println(1)
	fmt.Println("staring----")
	defer fmt.Println(2)
	fmt.Println("ending----")
	defer fmt.Println(3)
}
