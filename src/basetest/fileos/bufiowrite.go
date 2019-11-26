package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, e := os.OpenFile("./fulinhu.txt", os.O_WRONLY|os.O_CREATE, 7777)
	writer := bufio.NewWriter(file)
	defer file.Close()
	if e != nil {
		fmt.Println("文件与错误", e)
		return
	}
	for i := 0; i < 100; i++ {

		s := strconv.Itoa(i) + "fulinhu\n"
		fmt.Println(s)
		writer.WriteString(s)
	}
	writer.Flush()
}
