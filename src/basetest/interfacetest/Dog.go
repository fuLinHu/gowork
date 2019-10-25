package main

import "fmt"

type Dog interface {
	eat()
}

type fugu struct {
	name string
}
type gubin struct {
	name string
}

func (f fugu) eat() {
	fmt.Println(f.name)
}

func (g gubin) eat() {
	fmt.Println(g.name)
}
func main() {
	var dog Dog
	dog = &fugu{name: "fugu"}
	dog.eat()
	dog = gubin{name: "gubin"}
	dog.eat()
}
