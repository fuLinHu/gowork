package main

import "fmt"

func out(ch1 chan <- int)  {
	for i:=0;i<100;i++{
		fmt.Println("向通道接发送的数据为",i)
		ch1 <-i
	}
	close(ch1)
}

func in(ch <- chan int){
	for i:=range ch{
		fmt.Println("从通道接受到的数据为",i)
	}
}

func main(){
	ch:=make(chan int)
	go out(ch)

	in(ch)

}