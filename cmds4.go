package main

import (
	"fmt"
)

func main(){
	/*
	existem 3 tipos de variaçoes de print
	1)
	aquela que possui 
	-printf que na sua estrutura vai ter primeiro o texto nas aspas duplas e depoi vai ter uma virgula e as variaveis que serao encaixadas dentro das aspas
	fmt.Printf("%s is %d years old.\n", name, age)
	-prinln exibe  e pula uma linha
	fmt.Println(name, "is", age, "years old.")
	
	2)
	nesse caso vao existir as mesmas variacoes de Sprintln e Sprintf
	Então o fmt.Sprint() vai trandoramr tudo dentro dele em uma string, podendo fazer seu uso dessa forma
 	*/
	 x := 5
	 z := 4
 
	 a := fmt.Sprint(x,z)
	 fmt.Printf(a)
	/*
	atribuindo para um variavel para receber os valores em string
	*/
	/*
	
	*/
}