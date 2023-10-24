package main

import "fmt"

func main() {
	const name string = "Hari"
	const age int = 25

	if name == "Novrial" && age == 25 {
		fmt.Println("Ini Novrial")
	} else if name == "Sandi" {
		fmt.Println("Ini Sandi")
	} else {
		fmt.Println("Bukan Novrial Sandi")
	}

	for i := 1; i <= 5; i++ {
		star := ""
		for j := 1; j < i; j++ {
			star += "* "
		}
		fmt.Println(star)
	}
}
