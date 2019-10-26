package main

import (
	"reflect"
	"fmt"
)

func reflecttest(v interface{})  {
	of := reflect.ValueOf(v)

	if of.Elem().Kind()==reflect.Int{
		of.Elem().SetInt(15)
		fmt.Println(of.Elem())
	}

}
type stu struct {

}
func main() {
	//var i =12.0
	//reflecttest(&i)
	var j *int
	fmt.Println(reflect.ValueOf(j).IsNil())

	var st stu = stu{}
	of := reflect.ValueOf(st)
	fmt.Println("是否有值",of.FieldByName("abc").IsValid())
	//fmt.Println(of.IsNil())
	fmt.Println(of.IsValid())
}
/*func main() {
	// *int类型空指针
	var a *int
	fmt.Println("var a *int IsNil:", reflect.ValueOf(a).IsNil())
	// nil值
	fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())
	// 实例化一个匿名结构体
	b := struct{}{}
	// 尝试从结构体中查找"abc"字段
	fmt.Println("不存在的结构体成员:", reflect.ValueOf(b).FieldByName("abc").IsValid())
	// 尝试从结构体中查找"abc"方法
	fmt.Println("不存在的结构体方法:", reflect.ValueOf(b).MethodByName("abc").IsValid())
	// map
	c := map[string]int{}
	// 尝试从map中查找一个不存在的键
	fmt.Println("map中不存在的键：", reflect.ValueOf(c).MapIndex(reflect.ValueOf("娜扎")).IsValid())
}*/