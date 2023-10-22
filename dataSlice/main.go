package main

import "fmt"

func main() {
	var month = [...]string{
		"Januari",
		"Febuari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}
	var slice1 = month[4:7]

	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	month[5] = "Diubah"

	fmt.Println(slice1)
}
