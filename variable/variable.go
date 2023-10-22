package main

import "fmt"

func main() {
	var name string

	name = "Novrial"
	fmt.Println(name)

	var color = "blue"
	fmt.Println(color)

	var umur = 25
	fmt.Println(umur)

	// more simple

	city := "Yogyakarta"
	fmt.Println(city)

	// multiple declaration

	var (
		firstName = "Novrial"
		lastName  = "Sandi"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}
