package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	grandezza := os.Args[1]
	grnd, _ := strconv.Atoi(grandezza)
	x1 := os.Args[2]
	y1 := os.Args[3]
	x, _ := strconv.Atoi(x1)
	y, _ := strconv.Atoi(y1)
	var conta, cont int
	cont = 1
	for i := 0; i <= grnd; i++ {
		fmt.Print(conta)
		conta++
		if conta == grnd+1 {
			for j := 0; j < grnd; j++ {

				fmt.Print("\n", cont)
				cont++
				if cont == y+1 {
					for i := 0; i < x-1; i++ {
						fmt.Print(" ")
					}
					fmt.Print("*")
				}
			}
		}

	}

}
