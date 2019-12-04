package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func onces() {
	fmt.Println("运行的onces")
}

func onceed() {
	fmt.Println("运行了onced")
}

func test() {
	once.Do(onceed)
}

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println("运行了几次", i)
		go test()
	}
}
