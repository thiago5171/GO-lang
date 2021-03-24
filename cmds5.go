package main

import (
	"fmt"
)

// para criar um tipo basta fazer neste modelo abaixo,sendo int seu tipo subjacente
type hotdog int

//para atribuir o tipo a uma variavel
var b hotdog = 6

func main() {

	fmt.Printf("%T\n", b)
	fmt.Println(b)
	//nao é possivel fazer isso, porque estaria fazendo operaçoes ou atribuiçoes com tipos diferentes
	//soma = 5+b
	//fmt.Println(soma)
	// mas é possivel caso converta  o b para virar do mesmo tipo que o valor somado
	var soma = 5 + int(b)
	fmt.Println(soma)

	soma1 := fmt.Sprintf("%v", soma)
	// caso queira converter em bytes algum elemento  basta fazer o seguinte

	var bytes = []byte(soma1)
	fmt.Printf("quantidade de bytes usasdo foi =%v \n", bytes)

	// para exibir um texto da mesma forma que esta sendo digitada em um progrmaa basta utilizar ` texto desejado`
	s := `pipipip0i
		tste 
			texte2`
	fmt.Printf("%v",s)

}
