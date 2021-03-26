package main

import (
	"fmt"
)

func main(){
	x:=2
	switch{
	case (x==1):
		fmt.Println("x=1")
	case (x==2):
		fmt.Println("x=2")
	default:
		fmt.Println("erro!")
	}

}

//Crie um programa que utilize a declaração switch, sem switch statement especificado.
