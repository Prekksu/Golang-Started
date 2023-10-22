package main

import "fmt"

func main() {
	type NoKTP int
	type status bool

	var ktp NoKTP = 12312412
	var marriedStatus status = true

	fmt.Println(ktp)
	fmt.Println(marriedStatus)
}
