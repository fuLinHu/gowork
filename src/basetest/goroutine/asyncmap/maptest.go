package main

import (
	"strconv"
	"sync"
	"fmt"
)

func main() {
	sy:=sync.WaitGroup{}
	var  mp =make(map[string]int)
	for i:=0;i<5;i++{
		sy.Add(1)
		go func(n int){
			key:=strconv.Itoa(n)
			mp[key] = n
			fmt.Println("运行到第",i,"个")
			sy.Done()
		}(i)
	}
	sy.Wait()

}
