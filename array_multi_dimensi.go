package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Masukan Baris :")
	scanner.Scan()
	BarisInt, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Masukan Kolom :")
	scanner.Scan()
	KolomInt, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println(err)
	}

	hasil := make([][]uint8, BarisInt)
	for i := 0; i < BarisInt; i++ {
		hasil[i] = make([]uint8, KolomInt)
		for j := 0; j < KolomInt; j++ {
			hasil[i][j] = uint8(i*i + 1)
		}
	}
	fmt.Println(hasil)
}
