package main

import (
	"basetest/package/testone"
	"fmt"
)

func init() {
	fmt.Println("main init")
}
func main() {
	s := testone.Name
	fmt.Println(s)
}
