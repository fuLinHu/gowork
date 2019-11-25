package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, e := os.Open("C:\\Users\\12056\\Desktop\\问题.txt")
	if e != nil {
		fmt.Println("文件存在错误！！")
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		s, e := reader.ReadString('\n')
		if e == io.EOF {
			fmt.Println("文件读取完毕！！！")
			break
		}
		fmt.Println(s)
	}
}
