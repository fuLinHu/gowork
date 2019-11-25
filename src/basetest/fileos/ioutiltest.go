package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	bytes, e := ioutil.ReadFile("C:\\Users\\12056\\Desktop\\问题.txt")
	if e != nil {
		fmt.Println("文件有问题！！！")
	}
	fmt.Println(string(bytes))
}
