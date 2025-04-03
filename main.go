package main

import "fmt"

func main() {
var ages=[4]int{17,16,20,40}
	nomes:=[4]string{"Maria","gabi","ana","carol"}
	fmt.Println(ages)
	fmt.Println(nomes)
	// Slice 
	var score = []int{100,200,300,400}
	fmt.Println(score)
	score[1] = 2
	fmt.Println(score,len(score), cap(score))
	rangeOne := score [1:3]
	fmt.Println (rangeOne)
	rangeTwo := score[2:]
	fmt.Println(rangeTwo)
	rangeThree := score [:3]
	fmt.Println(rangeThree)

	var superherois =[]string{"Deadpool", "Homem-aranha","Motoqueiro fantasma"}
	fmt.Println(superherois)
superherois = append(superherois, "Ben 10")
fmt.Println(superherois, len(superherois),cap (superherois))
 }
 	
 
	

	
