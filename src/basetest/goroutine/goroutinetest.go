package main

import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup
func main() {
	go test()
	fmt.Println("main method run---")
	//time.Sleep(time.Second)

	for i:=0;i<10;i++{
		wg.Add(1)
		go hello(i)
	}
	wg.Wait()
}

func test(){
	fmt.Println("test method run---")
}
func hello(i int){
	defer wg.Done()
	fmt.Println("运行-----",i)
}

