package main

import (
	"fmt"
	"reflect"
)

func testRe(a interface{})  {
	of := reflect.TypeOf(a)
	name := of.Name()
	kind := of.Kind()
	fmt.Println("TypeOf---",of)
	fmt.Println("name---",name)
	fmt.Println("kind---",kind)
	fmt.Println("-------------valueOf----------------")
	valueOf := reflect.ValueOf(a)
	fmt.Println("valueOf---",valueOf)
	Int := reflect.Int
	fmt.Println("Int---",Int)

}

func getValue(va interface{}) interface{} {
	of := reflect.TypeOf(va)
	kind := of.Kind()
	valueOf := reflect.ValueOf(va)
	i := valueOf.Kind()
	fmt.Println(i)
	fmt.Println(valueOf)
	var inter interface{}
	switch kind {
	case reflect.Int:
		fmt.Println("Int",valueOf.Int())
		inter = valueOf.Int()
	case reflect.Int8:
		fmt.Println("Int8",int8(valueOf.Int()))
		inter = int8(valueOf.Int())
	case reflect.Int16:
		fmt.Println("Int16",int16(valueOf.Int()))
		inter = int16(valueOf.Int())
	case reflect.String:
		fmt.Println("String",string(valueOf.String()))
		inter = string(valueOf.String())
	}
	return inter
}

func main() {
	/*var i = "fulinhu"
	testRe(i)
	fmt.Println("------")
	type stu struct {

	}
	var s stu
	testRe(s)
	fmt.Println("------")

	var m map[int]string
	testRe(m)
	fmt.Println("------")*/
	var i =100
	value := getValue(i)
	fmt.Println("结果是：：",value)

}
