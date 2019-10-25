package main

import (
	"fmt"
	"strings"
)

func main() {

	var s = "fulinhu"
	i := len(s)
	index := strings.Index(s, "hu")
	fmt.Println(i, index)

}
