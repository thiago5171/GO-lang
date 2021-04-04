package main

import (
	"fmt"
)
// 'ARRAY' ele sempre comeca contar de 0 e sao imutaveis, uma vez predefinida  a quantidade de casas n se pode alterar
var x [5]int
var y [3]int{1,2,3}
func main(){
	x[0]=1
	x[1]=2

	fmt.Printf("%T",x)
	fmt.Println(x[1])
	fmt.Println(x[0])
//é possivel usar o comando 'len()' para descobrir qunatos casas tem o array
	fmt.Println("x tem esse total   de arrays ",len(x))
// ALEM DISSO,	é possivel usar essa estrutura para definir os valores de array
array:= [5]int{1,2,3,4,5}
fmt.Println(array)

fmt.Println(y)

}