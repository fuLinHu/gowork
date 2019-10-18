package main

func main() {

	var ptr *int

	var ptr1 *[]string
	println(ptr == nil)
	println(ptr1 == nil)
	var ptrArr [5]*int
	var arr [5]int
	for i := 0; i < 5; i++ {
		arr[i] = i * 10
		ptrArr[i] = &arr[i]
	}
	println("-----------------")
	for i := 0; i < 5; i++ {
		println(*ptrArr[i])
	}
	println("-----------------")
	for i := 0; i < 5; i++ {
		println(arr[i])
	}
	println("-----------------------------")
	var k int = 10
	var ptrr *int
	var ptrrr **int

	ptrr = &k
	ptrrr = &ptrr
	println(*ptrr)
	println(*ptrrr)
	println(**ptrrr)

}
