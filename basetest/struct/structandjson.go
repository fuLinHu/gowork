package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string `json:"name"`
	Age  int `json:"age"`
	Adre	`json:"adre"`
	Sore []int `json:"sore"`
}

type Adre struct {
	City    string
	Country string
}

func main() {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
	}
	student :=Student{}
	student.Name = "付林虎"
	student.Age = 20
	student.Sore = s
	a := Adre{
		City:    "北京",
		Country: "中国",
	}
	student.Adre = a
	fmt.Println(student)
	data, e := json.Marshal(student)
	fmt.Println(string(data))
	fmt.Println(e)
	var stu Student
	unmarshal := json.Unmarshal([]byte(data), &stu)
	fmt.Println(unmarshal)
	fmt.Println(stu)
}
