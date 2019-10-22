package main

import (
	"errors"
	"fmt"
)

type error interface {
	Error() string
}

func sqlr(f int) (string, error) {
	if f < 0 {
		return "fulinhusuccess", errors.New("math: square root of negative number")
	}

	return "fulinhu", nil
}
func main() {
	result, err := sqlr(-1)
	fmt.Println(result)
	if err != nil {
		fmt.Println(err)
	}
}
