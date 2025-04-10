package main
 
 import "fmt"
 
 func main() {
var saldo int
var opcao string
var deposito int
var saque int
fmt.Println("digite seu saldo")
fmt.Scan(&saldo)
for {
fmt.Println("voce quer 'sacar' ou 'depositar'?")
fmt.Scan(&opcao)
if opcao == "sacar"{
fmt.Println("quanto você deseja sacar?")
fmt.Scan(&saque)
fmt.Println("seu saldo é: ")
fmt.Scan("&saque")
saldo = saldo - saque
fmt.Println(saldo)
}
if opcao == "depositar"{
fmt.Println("quanto vc quer depositar?")
fmt.Scan(&deposito)
fmt.Println("seu saldo é:")
fmt.Scan("&deposito")
saldo = saldo + deposito
fmt.Println(saldo)
}
fmt.Println("Você deseja fazer mais uma transação? s/n")
var opcao1 string
fmt.Scan(&opcao1)
if opcao1 == "n" {
	break
}
}
}