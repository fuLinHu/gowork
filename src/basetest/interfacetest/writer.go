package main

import "fmt"

type Animal interface {
	run()
	eat()
}
type cat struct {
	name string
	age  int
}
type dog struct {
	name string
	age  int
	cat
}

func (c cat) run() {
	fmt.Println("this is cat run" + c.name)
}
func (d dog) run() {
	fmt.Println("this is cat run" + d.name)
}
func (d dog) eat() {
	fmt.Println("this is cat eat" + d.name)
}

/*func (c cat)eat()  {
	fmt.Println("this is cat eat"+c.name)
}*/

func main() {
	cat1 := cat{
		name: "cat",
	}
	dog1 := dog{
		name: "dog",
		cat: cat{
			name: "dogcat",
		},
	}
	/*	cat1.eat()*/
	cat1.run()
	dog1.run()
	dog1.cat.run()
	dog1.run()

}
