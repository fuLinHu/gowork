package main

import "fmt"

func main() {
	a := 10
	b := &a
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(b)
	fmt.Println(**&b)
	var c = 101

	sum(c)
	fmt.Println(c)

	fmt.Println("------------------")
	i := sum1(&c)
	fmt.Println(c)
	fmt.Println(i)

}

//传递的值
func sum(a int) int {
	a = 100
	return a
}

//传递的地址
func sum1(a *int) *int {
	*a = 1003
	return a
}
