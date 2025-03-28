package main

import "fmt"

func main() {
var admin string
var senha int
fmt.Print("Digite seu usuário: ")
fmt.Scan(&admin)
fmt.Println("o usuário é: ", admin)
fmt.Print("digite sua senha: ")
fmt.Scan(&senha)
fmt.Println("a senha é: ", senha)
if admin == "admin" &&  senha == 1234{
	fmt.Print("Acesso permitido! seja bem vindo")
}else {
	fmt.Printf("acesso negado, tente novamente")
}

 }
 	
 
	

	
