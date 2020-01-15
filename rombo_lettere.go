package main

import (
	"fmt"
)

func main() {
	var numero int
	fmt.Println("inserisci un numero")
	fmt.Scan(&numero)

	var stringa string = "A"
	var stringa2 string
	if numero < 0 {
		fmt.Println("numero negativo")
	} else {
		//Primo Ciclo for
		//PARTE SUPERIORE DEL ROMBO
		for i := 1; i <= numero; i++ {
			stringa2 = stringa
			for j := i; j <= numero; j++ {
				stringa2 = "  " /*servono due spazi per centrare il rombo*/ + stringa2 + "  "
				// il secondo "  " si occupa degli spazi di destra ma ee' solo un controllo in piu
			}
			fmt.Println(stringa2)
			if i != numero {
				//se l'indice del ciclo esterno e' diverso dal numero
				stringa = string(('A' + i)) + " " + stringa + " " + string('A'+i)
				//Il primo 'A'+i serve per andare avanti nell'alfabeto e scriverlo a destra
				//il secondo invece serve per fare la stessa cosa a sinistra
				//senza, il rombo, sarebbe formato solo dal carattere runa inserito
			}

		}
		//Secondo Ciclo for
		//PARTE INFERIORE DEL ROMBO
		for i := 0; i <= numero; i++ {
			if i < numero-1 {
				stringa2 = stringa[i*2+2 : len(stringa)-i*2-2] //serve per allineare il rombo
				for j := 0; j <= i+1; j++ {
					stringa2 = "  " /*servono due spazi per centrare il rombo*/ + stringa2 + "  "
				}
				fmt.Println(stringa2)

			}

		}
	}
}
