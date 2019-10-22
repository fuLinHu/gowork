package main

import "fmt"

func main() {
	p := person{
		18,
		"fulinhu",
		adress{
			city:    "beijing",
			country: "zhongguo",
		},
	}
	fmt.Println(p.string)
	fmt.Println(p.city)
}

type person struct {
	int
	string
	adress
}

type adress struct {
	city    string
	country string
}
