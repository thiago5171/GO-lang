package main

import (
	"fmt"
)

func main(){
	thiago := &Contato{
		Nome : "thiago",
		Telefone : "9999-9999",
		Email : "tgazaroli@gmail.com",
	}

	fmt.Printf(thiago)
}