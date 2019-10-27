package main

import (
	"reflect"
	"fmt"
)
type student struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Tall float32 `json:"tall"`
}

func(student)Getname(){
	fmt.Println("getname---------")
}

func (student)Getage()  {
	fmt.Println("getage---------")
}






func printallmethod(v interface{})  {
	of := reflect.ValueOf(v)
	method := of.NumMethod()
	fmt.Println(method)
	for i:=0;i<method;i++{
		var args = []reflect.Value{}
		of.Method(i).Call(args)
	}

}

func main() {
	st:=student{}
	printallmethod(st)
	/*type stu struct {
		name string
		age int
	}
	s:=stu{
		name:"fulinhu",
		age:10,
	}
	of := reflect.TypeOf(s)
	valueOf := reflect.ValueOf(s)

	field := of.NumField()
	for i:=0;i<field;i++{
		structField := of.Field(i)
		Name := structField.Name
		name, b := of.FieldByName(Name)
		byName := valueOf.FieldByName(Name)
		if(b){
			fmt.Println("----------------")
			fmt.Println("name is:",name.Name,byName)
			fmt.Println("----------------")
		}

		fmt.Println(Name,structField.Index,structField.Tag,structField.Type,structField.Anonymous)
	}*/



}
