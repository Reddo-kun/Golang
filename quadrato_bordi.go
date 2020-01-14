package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var m, n int

	M := os.Args[1]
	N := os.Args[2]
	m, _ = strconv.Atoi(M)
	n, _ = strconv.Atoi(N)

		//se n o m nulli

	if m<=0 || n<=0{		
		fmt.Println()
		return
	}
		//se n>=0.5*m

	if m<=2*n{			
		for o:=0;o<m;o++{
			for p:=0;p<m;p++{
				fmt.Print("*")
			}
			fmt.Println()
			}
			return
		}

		//struttura principale del quadrato

	for i:=0;i<m;i++{		
		if i<n||i>=(m-n){
			for j:=0;j<m;j++{
				fmt.Print("*")
			}
		}else{
			for k:=0;k<m;k++{
				if k<n||k>=(m-n){
					fmt.Print("*")
				}else{
					fmt.Print(" ")
				}
			}
		}
		fmt.Println()
	}
}
