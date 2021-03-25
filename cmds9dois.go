package main

import (
	"fmt"
)
// o comando 'inteface{}' permite que possamos atribuir qualquer valor a variavel, possibilitando implementar
//no switch um interpretador para saber qual o tipo da variavel e então executar uma ocasião
var x interface{}
func main(){
	x = true// altere os valores de x para ver as diferentes saidas
	switch x.(type){
	case int:
		fmt.Println("é um int")
	case float64:
		fmt.Println("é um float ")
	case bool:
		fmt.Println("é um boleano")
	case string:
		fmt.Println("é uma string")
 	}
/*
operadores logicos condicionais 
&& -->E
|| -->OU
*/

}