package main

import "fmt"

func main() {
	var name1 = "novrial"
	var name2 = "sandi"

	var result bool = name1 == name2
	fmt.Println(result)

	var val1 = 100
	var val2 = 50

	fmt.Println(val1 > val2)
	fmt.Println(val1 < val2)
	fmt.Println(val1 == val2)
	fmt.Println(val1 != val2)
}
