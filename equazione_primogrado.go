package main

import "fmt"

func main() {
	var a, b int
	fmt.Println("inserisci il valore di a e di b")
	fmt.Scan(&a, &b)
	fmt.Println("equazione : ", a, "x +", b, " = 0")
	equazione(a, b)
	//fmt.Println("x =", equazione(a, b))
}
func equazione(a1 int, b1 int) {
	var risultato int
	if b1 < a1 {
		fmt.Println("x =", -b1, "/", a1)
	} else {
		risultato = -b1 / a1
		fmt.Println("x =", risultato)
		//return risultato
	}

}
