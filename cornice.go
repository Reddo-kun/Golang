package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	M := os.Args[1]
	N := os.Args[2]
	m, _ := strconv.Atoi(M)
	n, _ := strconv.Atoi(N)
	for i := 0; i < m; i++ {
		fmt.Print("+")
		fmt.Print(strings.Repeat("=", n))
	}
	fmt.Print("+")
}
