package main

import "fmt"

func main() {
	var num1, num2, maxDiff, i, N int
	maxDiff = 0
	fmt.Println("Quanti numerei vuoi inserire?")
	fmt.Scan(&N)
	fmt.Println("inserisci il primo numero")
	fmt.Scan(&num1)

	for i = 1; i < N; i++ {
		fmt.Println("inserisci un altro numero")
		fmt.Scan(&num2)

		differenza := diff(num2, num1)
		if differenza > maxDiff {
			maxDiff = differenza
		}
		num1 = num2
	}
	fmt.Println("differenza massima tra due numeri consecutivi:", maxDiff)
}
func diff(num2 int, num1 int) int {
	var difference int
	if num2 > num1 {
		difference = num2 - num1
	} else {
		difference = num1 - num2
	}
	return difference
}
