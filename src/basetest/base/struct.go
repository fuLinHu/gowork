package base

import "fmt"

type book struct {
	name string
	age  int
}

func main() {

	var Book book
	Book.age = 10
	Book.name = "flh"
	///Book.bir=dwarf.Data{}
	println(Book.name)
	test1(Book)
	fmt.Println(Book)
}

func test1(param book) {
	println(param.name)
	println(param.age)

}
