package main

import "fmt"

func main() {
	var x interface{}
	x = "fulinhu"

	/*s,ok := x.(int)
	fmt.Println(s)
	fmt.Println(ok)*/
	judge(x)
}

func judge(typ interface{}) {
	switch x := typ.(type) {
	case string:
		fmt.Println("string", x)
	case int:
		fmt.Println("int", x)
	case float32:
		fmt.Println("float32", x)
	case bool:
		fmt.Println("bool", x)
	default:
		fmt.Println("other", x)

	}

}
