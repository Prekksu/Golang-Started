package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Novrial"
	names[1] = "Hari"
	names[2] = "Sandi"

	fmt.Println(names)
	fmt.Println(names[2])

	var values = [3]int{
		90, 80, 85,
	}
	fmt.Println(values)
	fmt.Println(len(values))
	fmt.Println(values[1])

}
