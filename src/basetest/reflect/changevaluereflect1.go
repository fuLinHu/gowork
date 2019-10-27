package main

import (
	"reflect"
	"fmt"
)

func main() {
	m:=map[string]int{}
	m["你好"]=10
	of := reflect.ValueOf(m)
	index := of.MapIndex(reflect.ValueOf("你好"))
	fmt.Println(index)
	valid := index.IsValid()
	fmt.Println(valid)
}