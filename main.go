package main
 
 import (
	"fmt")


	var valorsaldo int
	var opcao string
	var valordeposito int
	var valorsaque int

func Saldo() {
		fmt.Println("digite seu saldo")
 		fmt.Scan(&valorsaldo)
}
func Opcao(){
	fmt.Println("voce quer 'sacar' ou 'depositar'?")
	fmt.Scan(&opcao)
}
func Saque() {
	
		if opcao == "sacar"{
		fmt.Println("quanto você deseja sacar?")
		fmt.Scan(&valorsaque)
		valorsaldo = valorsaldo - valorsaque
		fmt.Println("seu saldo é: ", valorsaldo)
		
   }
}
func Deposito() {
	if opcao == "depositar"{
		fmt.Println("quanto vc quer depositar?")
		fmt.Scan(&valordeposito)
		valorsaldo = valorsaldo + valordeposito

		fmt.Println("seu saldo é:", valorsaldo)
	
		}
}
		
func main (){
Saldo()
Opcao()
Saque()
Deposito()
}
