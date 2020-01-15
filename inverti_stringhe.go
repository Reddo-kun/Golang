package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("inserisci stringa base")
	base := ""
	scannerB := bufio.NewScanner(os.Stdin)
	if scannerB.Scan() {
		base = scannerB.Text()
		base += "\n"
	}
	fmt.Println("insersci inverti")
	inverti := ""
	scannerI := bufio.NewScanner(os.Stdin)
	if scannerI.Scan() {
		inverti = scannerI.Text()
		inverti += "\n"
	}
	frase := strings.Fields(base) //crea una slice della stringa base
	for i := 0; i < len(frase); i++ {
		if strings.Contains(inverti, frase[i]) {
			fmt.Print(inv(frase[i]) + " ")
		} else {
			fmt.Print(frase[i] + " ")
		}
	}
}
func inv(s string) string {
	var s2 string
	runes := []rune(s)
	for i := len(runes) - 1; i >= 0; i-- {
		s2 += string(runes[i])
	}
	return s2
}
