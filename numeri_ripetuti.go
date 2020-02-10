package main

import (
	"fmt"
)

func unique(intSlice [5]int) int /* []int*/ {
	keys := make(map[int]bool)
	//list := []int{}
	var cont, nessuno int

	nessuno = 0
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {

			keys[entry] = true
			/*
				list = append(list, entry)
			*/
		} else if cont == 0 {
			//	fmt.Println("Ci sono numeri doppi")
			cont++
			nessuno = 2

		}

	}
	if nessuno == 0 {
		//fmt.Println("Non ci sono numeri doppi")

	}
	//return list
	return nessuno
}

func main() {
	var list [5]int
	var i, zero int
	for {
		fmt.Println("inserisci i numeri")
		for i = 0; i < 5; i++ {
			fmt.Scan(&list[i])
			if list[i] == 0 {
				zero = 3
				break

			}
		}

		if zero == 3 {
			break
		} else {

			valore := unique(list)
			if valore == 0 {
				fmt.Println(list)
				fmt.Println("Non ci sono numeri doppi")
			} else if valore == 2 {
				fmt.Println(list)
				fmt.Println("Ci sono numeri doppi")
			}
		}
	}

	//intSlice := []int{1, 5, 3, 8}
	//uniqueSlice :=
	//	fmt.Println(uniqueSlice)

}
