package main

import "fmt"

func main() {
	a, b :=10, 3
	fmt.Println("a soma é: ", a + b)
	fmt.Println("A multiplicação é: ", a * b)
	fmt.Println("A divisão é: ",a / b)
	fmt.Println("subtração é  ", a - b)
	fmt.Println("O resto da divisão é ",a % b)
	a++
	fmt.Println("Incrementar a", a)
	if a >0 && b>0 {
		fmt.Println("numeros positivos")
	}
}