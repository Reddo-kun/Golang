package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var num int
	numero := os.Args[1]

	num, _ = strconv.Atoi(numero)

	for i := 0; i < num; i++ {
		for j := 0; j < 2*num+1; j++ {
			if j == num && i == 0 {
				fmt.Print("+")
			} else if j <= num+i && j >= num-i {
				fmt.Print(" ")
			} else {
				fmt.Print("0")
			}

		}
		fmt.Println()
	}

}
