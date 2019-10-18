package main

import (
	"fmt"
	"strconv"
)

func main() {
	var mapt = map[string]string{}
	mapt["你好"] = "hello"
	mapt["名称"] = "name"
	kk, jj := mapt["名称1"]
	fmt.Println(kk)
	fmt.Println(jj)
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			mapt[strconv.Itoa(i)] = strconv.Itoa(i) + "你好吗"
		}
	}

	println("map的长度是：", len(mapt))

	for i := 0; i < 100; i++ {
		_, ok := mapt[strconv.Itoa(i)]
		println(ok)
		if ok {
			println("存在这样的k", i)
			println(mapt[strconv.Itoa(i)])
			delete(mapt, strconv.Itoa(i))
			_, ok1 := mapt[strconv.Itoa(i)]
			if ok1 == false {
				println("删除了", i)
			}
		} else {
			println("---------------------------")
			println("不存在这样的k", i)
		}

	}

}
