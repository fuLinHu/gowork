package base

func main() {
	var arr []int
	for i := 0; i < 10; i++ {
		arr = append(arr, i*10)
	}
	for i, j := range arr {
		println(i)
		println(j)
		println("------------")
	}
}
