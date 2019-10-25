package base

import "fmt"

func main() {
	var arr = make([]int, 3, 9)
	arr = append(arr, 1)
	arr = append(arr, 1)
	arr = append(arr, 1)
	arr = append(arr, 1)
	arr = append(arr, 1)
	arr = append(arr, 1)
	arr = append(arr, 1)
	arr = append(arr, 1)
	arr = append(arr, 1)
	arr = append(arr, 1)
	arr = append(arr, 1)
	arr = append(arr, 1)
	arr = append(arr, 1)
	arr = append(arr, 1)
	arr = append(arr, 1)
	arr = append(arr, 1)
	fmt.Println(arr[2:6])
	fmt.Println(len(arr))
	fmt.Println(cap(arr))
	arry := []int{}
	var array []int
	println(len(arry))
	println(&arry)
	if array == nil {
		fmt.Println("空切片")
	}
	arry = append(arry, 2)
	fmt.Println(arry)

}
