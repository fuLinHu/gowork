package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//C:\\Users\\12056\\Desktop\\问题.txt
	file, e := os.Open("C:\\Users\\12056\\Desktop\\excel（智能油田）\\对象-采集201910.xls(1).xlsx")
	if e != nil {
		fmt.Println("open file failed!, err:", e)
		return
	}
	defer file.Close()
	if e == io.EOF {
		fmt.Println("文件读取完毕！")
		return
	}
	if e != nil {
		fmt.Println("open file failed!, err:", e)
		return
	}

	var all []byte
	tem := make([]byte, 128)
	for {
		n, err := file.Read(tem)
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		fmt.Printf("读取了%d字节数据\n", n)
		all = append(all, tem[:n]...)
	}
	fmt.Println(string(all))

}
