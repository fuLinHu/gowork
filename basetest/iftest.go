package main

import "fmt"

func main() {
	/*var i = 90;
	if(i<200){
		println("你好！")
	}else{
		println("我不好好！")
	}

	if(i<60){
		println("不及格")
	}else if(i>=60&&i<80){
		println("及格")
	}else if(i>=80&&i<90){
		println("良好")
	}else{
		println("优秀")
	}

	var B = 100
	switch B {
		case 10:
			fmt.Println("是10")
		case 20:
			fmt.Println("是20")
		case 30:
			fmt.Println("是30")
		default:
			fmt.Println("无论如何都要运行的")

	}*/
	var x interface {
	}
	x = 100
	switch j := x.(type) {
	case int:
		fmt.Printf("x 是 int 型", j)
	case bool:
		fmt.Printf("x 是 bool 型", j)
	case float32:
		fmt.Printf("x 是 bool 型", j)
	}
}
