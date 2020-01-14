package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var num int
	numero := os.Args[1]                   //prende da linea di comando e lo inserisce nell'argomento[1]
	carattere := os.Args[2]
	num, _ = strconv.Atoi(numero)         //da stringa a intero "strconv"
	if num > 0 {
		for i := 0; i < num; i++ {               
			for j := 0; j < 2*num; j++ {          
				if j <= num+i && j >= num-i {   //if usato per stampare l'albero
					fmt.Print(carattere)
				} else {                          //else invece viene usato per stampare gli spazi
					fmt.Print(" ")
				}
			}
			fmt.Println()              //Println usato per andare a capo
		}
	} else {
		fmt.Println("numero troppo piccolo")
	}
}
