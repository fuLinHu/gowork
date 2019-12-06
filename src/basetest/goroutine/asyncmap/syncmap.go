package main

import (
	"sync"
	"strconv"
	"fmt"
)

var m = sync.Map{}
var sy = sync.WaitGroup{}
func main() {
	for i:=0;i<10000;i++{
		sy.Add(1)
		func(n int){
			key:=strconv.Itoa(i)
			m.Store(key,n)
			fmt.Println("运行到第",i,"个")
			defer sy.Done()
		}(i)
	}
	sy.Wait()
}
