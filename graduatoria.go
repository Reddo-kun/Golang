package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	str := ""
	var stringa []string
	fmt.Println("Scrivi nome e punteggio")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if input == "" {
			//lettura terminata
			break
		} else {
			str += input + "\n"
			stringa = strings.Fields(str)
		}
	}
	var nome, punt string
	max := ""
	nomevinc := ""
	for i := 0; i < len(stringa)-1; i = i + 2 {
		nome = stringa[i]
		punt = stringa[i+1]
		fmt.Println(nome, punt)
		if punt >= max {
			max = punt
			nomevinc = nome
		}
	}
	fmt.Println("punteggio massimo:", max, "di", nomevinc)
}
