package base

import "fmt"

type phone interface {
	call()
}

type Iphone struct {
}

type Huawei struct {
}

func (iphone Iphone) call() {
	fmt.Println("运行的是苹果机")
}

func (huawei Huawei) call() {
	fmt.Println("运行的是华为")
}

func main() {
	var ph phone
	ph = new(Iphone)
	ph.call()

	ph = new(Huawei)
	ph.call()
}
