package testone

import (
	"basetest/package/testtwo"
	"fmt"
)

//测试小小写
var Name = "fulinhu"

var age = 20

func init() {
	fmt.Println(Name)
	fmt.Println("这个是一个inint--testone-")
	testtwo.Tworun()
}
