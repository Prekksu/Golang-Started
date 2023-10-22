package main

import "fmt"

func main() {
	var a = 10
	var b = a + 10

	fmt.Println(b)

	// augmented assigments
	var i = 20
	i += 10
	fmt.Println(i)

	// unary operator
	var j = 10
	j++ // j + 1
	fmt.Println(j)
}
