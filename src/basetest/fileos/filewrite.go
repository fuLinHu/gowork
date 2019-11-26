package main

import (
	"fmt"
	"os"
)

func main() {
	file, e := os.OpenFile(".//a.txt", os.O_CREATE|os.O_RDWR, 7777)
	if e != nil {
		fmt.Println("文件错误")
		return
	}
	b := []byte("付林虎\n")
	file.Write(b)
	file.WriteString("hello 小王子") //直接写入字符串数据
	defer file.Close()
}
