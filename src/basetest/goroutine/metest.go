package main

import (
	"fmt"
)

func run() {
	for i := 0; i < 100; i++ {
		fmt.Println(i, "运行到。。。。。。")
	}
}

//生产者
func produce(ch chan<- int, ch2 chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("生产了", i)
		/*ch2 <- i*/
	}
}

//消费者
func custom(ch <-chan int) {
	for i := range ch {
		fmt.Println("消费了", i)
	}
}

func main() {
	result := make(chan int, 5)
	ch := make(chan int)
	//无缓冲通道  接收和消费是同步的  只有在接收值的时候才有能去发送值 否则会因为发送的值不能被接收而发生阻塞 而发生死锁
	//因此起动一个线程去接收值，然后在去发送数据
	go custom(ch)
	produce(ch, result)
	/*for i:=0;i<5;i++{
		<-result
	}*/

}
