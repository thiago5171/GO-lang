package main

import (
	"fmt"
)

// para criar um tipo basta fazer neste modelo abaixo,sendo int seu tipo subjacente
type hotdog int
//para atribuir o tipo a uma variavel
var b hotdog =6

func main(){

	fmt.Printf("%T\n",b)
	fmt.Println(b)
	//nao é possivel fazer isso, porque estaria fazendo operaçoes ou atribuiçoes com tipos diferentes
	//soma = 5+b
	//fmt.Println(soma)
	// mas é possivel caso converta  o b para virar do mesmo tipo que o valor somado
	var soma = 5 + int(b)
	fmt.Println(soma)
}