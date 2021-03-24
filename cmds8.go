package main

import (
	"fmt"
)

func main(){
// Uso da estrutura IF
x:=10
// esse iff se lê " sera que essa sentença é vdd"
	if x < 100 {// se a sentença  for true entao executa a ação abaixo
		fmt.Println("x é maior que 100")
	}
// ou tambem podemos um if com a negação da sentença
// esse iff se lê " sera que essa sentença nao é vdd"
	if !( x > 100 ) { //o  '!' é o sinal de negação da sentença
		fmt.Println("x é menor que 100")
	}
// EXTENSÕES DO IF
// ELSE
	if x>5 {
		fmt.Println("x>5")
	}else{
		fmt.Println("x<5")
		}
	
	// ELSE IF
	if (x > 15){
		fmt.Println("x é maior que 15")	
	}else if (x < 15){
		fmt.Println("x é menor que 15")
	}else{
		fmt.Println("x é igual que 15")
	}

//ainda é possivel fazer uma declaração de uma variavel dentro do IF usando o ';'
/*
if x:=100; x<90000{
	
}
*/
}
