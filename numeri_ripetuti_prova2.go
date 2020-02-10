package main

import (
	"fmt"
	"os"
	"strconv"
)

func unique(intSlice []int) int /* []int*/ {
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
	//var list [5]int
	var lista []int
	for {
		list := os.Args[1:]
		//	fmt.Println("inserisci i numeri")
		/*	for {
			fmt.Scan(&list[i])
			i++
			if list[i] == 0 {
				zero = 3
				break

			} else {
				zero = 0
				break
			}

		}*/
		for _, valore := range list {
			val, _ := strconv.Atoi(valore)
			if val == 0 {
				break
			} else {

				lista = append(lista, val)
			}

		}

		/*	if zero == 3 {
			break
		} else {*/

		valore := unique(lista)
		if valore == 0 {
			fmt.Println(lista)
			fmt.Println("Non ci sono numeri doppi")
		} else if valore == 2 {
			fmt.Println(lista)
			fmt.Println("Ci sono numeri doppi")
		}
		break

		//}
	}

	//intSlice := []int{1, 5, 3, 8}
	//uniqueSlice :=
	//	fmt.Println(uniqueSlice)

}
