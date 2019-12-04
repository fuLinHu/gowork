package main

import "fmt"

func main() {

	var mymap = make(map[string]int)
	mymap["f"] = 1
	mymap["u"] = 2
	mymap["l"] = 3
	mymap["i"] = 4
	mymap["n"] = 5
	mymap["h"] = 6
	mymap["u"] = 7
	for key, v := range mymap {
		fmt.Println(key)
		fmt.Println(v)
	}
}
