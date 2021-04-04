package main

import "fmt"

func main() {

	for N := 1; N <= 100; N++ {
		if N%2 == 0 {
			fmt.Println(N, "Genap")
		} else {
			fmt.Println(N, "Ganjil")
		}
	}

}
