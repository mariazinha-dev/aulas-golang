package main
 
 import (
	"fmt")


func main (){
	alunoIdade := make(map[string]int)
	alunoIdade["carol"] =15 
	alunoIdade["maria"] = 15
	alunoIdade["aninha"] = 16
	alunoIdade["Matheus"] = 16
	fmt.Println("idade da carolzinha",alunoIdade["carol"])
	notasAlunos := map[string]float64{
		"carol" : 5.5,
		"maria" : 10,
		"aninha" : 7.5,
		"matheus":7.5,
	}
	for nome,nota := range notasAlunos{
	fmt.Printf("%s tirou a nota %f\n",nome,nota)
	}
}