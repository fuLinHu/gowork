package main

import (
	"sync"
	"fmt"
)

var x int
var wg sync.WaitGroup
var lock sync.Mutex
func sum(){
	for i:=0;i<5000;i++{
		lock.Lock()
		x+=i
		lock.Unlock()
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go sum()
	go sum()
	wg.Wait()
	fmt.Println(x)
}
