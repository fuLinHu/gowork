package main

import "fmt"

type DogP interface {
	eatp()
}

type fugup struct {
	name string
}
type gubinp struct {
	name string
}

func (f *fugup) eatp() {
	fmt.Println(f.name)
}

func (g *gubinp) eatp() {
	fmt.Println(g.name)
}
func main() {
	var dog DogP
	dog = &fugup{name: "fugu"}
	dog.eatp()
	dog = &gubinp{name: "gubin"}
	dog.eatp()
}
