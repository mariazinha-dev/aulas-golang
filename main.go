package main

import (
	"fmt"
)

func main () {
	estoque:= make(map[string]int)
	 estoque["Coxinha"] = 10
	 estoque["Pão de Queijo"] = 15
	 estoque["Refrigerante"] = 20
	 estoque["Coxinha"] -= 2
	 estoque["Pão de Queijo"] -= 1
	 estoque["Refrigerante"] -= 0

	 for estoque, total := range estoque{
		fmt.Printf("%s tem: \n %d quantidade \n", estoque, total)
	 }
}



	

 


 