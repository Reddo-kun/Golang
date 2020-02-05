package main

import "fmt"

var num, nume float64
var n int

func main() {
	var numero, cifre float64
	var cifredecimali int

	fmt.Println("inserisci un numero in float64")
	fmt.Scan(&numero)
	fmt.Println("inserisci numero cifre decimali")
	fmt.Scan(&cifredecimali)

	/*Condizione, se cifredecimali e' 0 allora il programma trasforma
	in int il numero inserito e lo stampa dopodiche riprende il numero
	in float64 e aggiunge 0.5, lo ritrasformma in int e lo stampa */

	if cifredecimali == 0 {
		n := int(numero)
		fmt.Println("troncamento:", n)
		numero = numero + 0.5
		narr := int(numero)
		fmt.Println("arrotodamento:", narr)

	} else {
		//Se cifredecimali e' diverso da zero esegue tutto il resto
		// ciclo per creare il numero in base 10 con cui si moltiplichera' o dividera' il float64
		for i := 0; i < cifredecimali; i++ {
			if i == 0 {
				cifre = 10
			} else {
				cifre = cifre * 10
			}

		}

		fmt.Println("troncamento:", troncamento(cifre, numero))      //Passo le due variabile alla funzione troncamento
		fmt.Println("arrotodamento:", arrotondamento(cifre, numero)) //Passo le due variabili alla funzione arrotondamento

	}

}
func arrotondamento(decimali float64, num float64) float64 {
	num = num * decimali //decimali e' il numero inserito dall'utente 10^n (2 cifre = 100 ecc)
	num = num + 0.5      //Sommo 0.5 per fare l'arrotondamento
	n = int(num)
	nume = float64(n)
	nume = nume / decimali //divido per 10^n
	return nume            //Passo al main il valore finale per farlo stampare
}

func troncamento(decimali float64, num float64) float64 {
	var numx100, nfloat, nd100 float64
	var nint int
	//Stessa cosa della func arrotondamento
	numx100 = num * decimali
	nint = int(numx100)
	nfloat = float64(nint)
	nd100 = nfloat / decimali
	return nd100

}
