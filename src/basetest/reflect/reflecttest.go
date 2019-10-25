package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i = "fulinhu"
	of := reflect.TypeOf(i)
	fmt.Println(of)
}
