package base

import "fmt"

func main() {
	var j = 0
	for i := 0; i < 100; i++ {
		j += i
	}
	fmt.Println(j)

	var kj = 0
	for kj < 10 {
		kj += 1
		fmt.Println("-----------", kj)
	}
}
