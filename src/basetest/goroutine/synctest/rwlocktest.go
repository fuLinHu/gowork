package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	y      int64
	wgy    sync.WaitGroup
	rwlock sync.RWMutex
	locky  sync.Mutex
)

func read() {
	rwlock.RLock()
	time.Sleep(time.Millisecond)
	fmt.Println(y)
	rwlock.RUnlock()
	wgy.Done()
}

func write() {
	//rwlock.Lock()
	time.Sleep(2 * time.Millisecond)
	y += 1
	//rwlock.Unlock()
	wgy.Done()
}

func main() {
	for i := 0; i < 10; i++ {
		wgy.Add(1)
		go write()
	}

	for i := 0; i < 1000; i++ {
		wgy.Add(1)
		go read()
	}
	wgy.Wait()
	fmt.Println("最后的结果为：：", y)

}
