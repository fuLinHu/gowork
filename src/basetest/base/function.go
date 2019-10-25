package base

type hand func(j, k int) int

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

/*func main() {
	var jj=max(3,4)
	println(jj)
	a,b:=swp("nihao","buhao")
	println(a,b)
	var f = func(i,j int)(int){
		return i-j
	}
	test(f)
}*/

func test(kk hand) {
	param1, param2 := 23, 45
	var result = kk(param1, param2)
	println(result)
}
func swp(x, y string) (string, string) {
	return y, x
}

func fibonacci(a int) int {
	println("--------------", a)
	if a < 2 {
		return a
	}
	return fibonacci(a-2) + fibonacci(a-1)
}

func main() {
	for i := 0; i < 10; i++ {
		println(fibonacci(i))
	}

}
