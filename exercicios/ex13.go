package main

import (
	"fmt"
)

func main(){
	timeFavorito:= "sao paulo"

	switch timeFavorito {
	case "cruzeiro":
		fmt.Println("seu esporte favorto é cruzeiro")
	case "flamengo":
		fmt.Println("seu esporte favorto é flamengo")
	case "sao paulo":	
	fmt.Println("seu esporte favorto é sao paulo")
	default:
		fmt.Println("nao encontrado")

	}

}
//Crie um programa que utilize a declaração switch, onde o switch statement seja uma variável do tipo 
//string com identificador "timeFavorito".