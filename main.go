package main

import (
	"fmt"
)

func mediaNotas(n1, n2 float64) (float64, string){

	media := (n1 + n2) / 2

	if media > 6 {
		return media, "aprovado"
	} else
	{
		return media, "reprovado"
	}
}
func main() {
	media, resultado := mediaNotas(9.5, 8.8)
	fmt.Println("Média:",media)
	fmt.Println("Resultado:", resultado)
}



	

 


 