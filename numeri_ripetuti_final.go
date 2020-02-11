package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	var line string

	for {
		fmt.Println("inserisci i numeri")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			line = scanner.Text()
		}
		result := strings.Split(line, " ")
		for _, valore := range result {

			val, _ := strconv.Atoi(string(valore))
			if val == 0 {
				os.Exit(0)
			}
			list = append(list, val)

		}
		stampa(list)
		list = nil
	}

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
