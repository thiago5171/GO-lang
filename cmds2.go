package main

import "fmt"

func main() {
	/* Declarações são diferentes  de atribuições, pois na declaração  a variavel analisa qual tipo primitivo 
	esta sendo declarado e recebe o mesmo tipo primitivo  do valor declarado
	usando este  comndao para fazer declações 
			:=
	a grosso modo seria uma tipagem automatica
	apos feito uma declaração em uma variavel, nao podera fazer novas declarações, apenas  atribuições
	por exemplo eu fiz isso
	a:= 5
	nao sera possivel que eu use novamente o := com o a novamnete
	a:=20
	apenas poderei usar uma atribuição normal
	a = 20
	no entano ainda podera fazer outra nessa variavel caso haja uma nova variavel, por exemplo 
	a,z := 20,60
	*/
	x :=  "g"
	y :=  222
	z := true
	fmt.Println(x,"<-um %t texto\n",
	y,"<- um valor numerico\n",
	z,"<- um valor boleano [true/false]\n")
	fmt.Printf("x: %j ",x)
}
