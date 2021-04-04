package main

import "fmt"

func main() {
	a := 0
	b := 1
	c := 0

	for N := 1; N <= 100; N++ {
		if N == 1 {
			fmt.Println(a)
		}
		if N == 2 {
			fmt.Println(b)
		}
		c = a + b
		a = b
		b = c
		fmt.Println(c)
	}
}
