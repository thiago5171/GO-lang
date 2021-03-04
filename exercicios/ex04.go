package main

import (
	"fmt"
)
type inteiro int
var x  inteiro
func main(){
	fmt.Printf("%v\n%T\n",x,x)
	x = 42
	fmt.Printf("%v\n",x)
	
}

