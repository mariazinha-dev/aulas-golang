package main

import (
"fmt"
"strings"
"sort"
)
func main () {
greeting := "Hello my friends"
fmt.Println(strings.Contains(greeting,"dog"))
fmt.Println(strings.ReplaceAll(greeting,"my", "mine"))
fmt.Println(strings.ToUpper(greeting))
fmt.Println(strings.Index(greeting,"my"))
ages:=[]int {50,80,10,}
sort.Ints(ages)
fmt.Println(ages)
index := sort.SearchInts(ages,50)
fmt.Println(index)
names:=[]string {"Carol","Maicon", "ana"}
sort.Strings(names)
fmt.Println(names)
fmt.Println(sort.SearchStrings(names,"Carol"))
x:=15

for x<5 {
	fmt.Println(x)
	x++
}

for i:=0; i <5; i++{
	fmt.Println("for 2: ", i)
}

for i:=0;  i<len(names); i++ {
	fmt.Println(names[i])
}

for index, value := range names{
fmt.Println("o índice é: ", index, "e o valor é: ", value)
}

for index, value := range ages{
fmt.Println("o índice é: ", index, "e o valor é: ", value)
}
}