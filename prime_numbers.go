package main

import "fmt"

func main() {
	for N := 1; N <= 100; N++ {
		i := 0
		for a := 1; a <= 100; a++ {
			if N%a == 0 {
				i++
			}
		}
		if (i == 2) && (N != 1) {
			fmt.Println(N)
		}
	}
}
