package main

import (
	"fmt"
)
type inteiro int
var x  inteiro
var y int
func main(){
	fmt.Printf("a variaval x tem valor = %v\ne tipo subjancente =%T\n",x,x)
	x = 42
	fmt.Printf("%v\n",x)
	// METODO PARA CONVERTER UM TIPO int()
	y = int(x)
	fmt.Printf("a variaval y tem valor = %v\ne tipo  =%T\n",y,y)

}

