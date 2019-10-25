package main

import "fmt"

type Animal struct {
	name string
}

func (a Animal) move() {
	fmt.Println(a.name, "----在运动")
}

func (a Animal) wang() {
	fmt.Println(a.name, "____在汪汪")
}

type dog struct {
	feet int
	Animal
}

type cat struct {
	feet int
	Animal
}

func main() {
	d := dog{
		feet: 8,
		Animal: Animal{
			name: "kimidog",
		},
	}
	d.wang()
	d.move()

	c := cat{
		feet: 4,
		Animal: Animal{
			name: "kimicat",
		},
	}
	c.move()
	c.wang()
}
