package main
 
 import (
	"fmt")

func dividir(dividendo int, divisor int) (int, string){
if divisor == 0 {
	return 0, "Erro na divisão por zero"
}
return dividendo/divisor, "Sem erro"
}

func main (){
resultado, erro := dividir(10,0)
if erro != "Sem erro" {
fmt.Println(erro)
} else{
	fmt.Println(" O resultado da divisão é:",resultado, erro)
}

soma,mult,sub := operaçãobasica(10,2)
fmt.Println(soma)
fmt.Println(mult)
fmt.Println(sub)


}
func operaçãobasica(a int, b int) (int, int, int){
	soma := a +b
	multiplicacao:= a*b
	subtracao := a-b
	return soma,multiplicacao,subtracao

}