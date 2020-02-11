package main

import (
	"fmt"
	"os"
)

func main() {

	for {
		inserimento()
	}

}
func control(intSlice []int) int {
	keys := make(map[int]bool)
	var cont, nessuno int
	nessuno = 0
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {

			keys[entry] = true

		} else if cont == 0 {
			cont++
			nessuno = 2

		}

	}
	return nessuno
}

func inserimento() {
	var list []int
	var v int
	fmt.Println("inserisci i numeri")
	for {
		fmt.Scan(&v)

		if v == 0 {
			os.Exit(0)
		} else if v == 99 {
			break
		} else {
			list = append(list, v)
		}

	}

	stampa(list)

	list = nil
}
func stampa(lista []int) {
	valore := control(lista)
	if valore == 0 {
		fmt.Println(lista)
		fmt.Println("Non ci sono numeri doppi")
	} else if valore == 2 {
		fmt.Println(lista)
		fmt.Println("Ci sono numeri doppi")
	}
}
