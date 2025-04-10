package main
 
 import "fmt"
 
 func main() {
	var idade int
	fmt.Println("digite sua idade: ")
	fmt.Scan(&idade)
	if idade <18 {
		fmt.Println("você é de menor, não pode entrar!")
	} else if idade >=18 && idade <=60{
		fmt.Println("você é de maior, pode trar meu rapaz/senhorita/monsieur")
	} else {
		fmt.Println("você é um idoso!")
	}
 }