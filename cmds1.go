/*
para que serve o package main no inicio do codigo
	ele orienta que este seja o codigo principal caso tenha varias pastas com codigos dao continuidade para
	este que é p principal
*/
package main

import (
	"fmt"
)

func () { // essa func main mostra onde sera o inicio da execução do codigo ate o fim da ultima chaves
	fmt.Println("heelo world", "ola pessoas", "demontrar numeros 100 e ", 100)

	//é possivel descobrir quantos bytes esta sendo usada determinada funcao e  o numeros de erros da seguinte forma,
	//porem é necessario exibilas depois
	numerodebytes, erros := fmt.Println("heelo world", "ola pessoas", "demontrar numeros 100 e ", 100)
	fmt.Println(numerodebytes, erros)
	//a fmt println retorna 2 valores, ou seja eh nescesario 2 variaveis  para  colocar 2 valores, mas caso eu nao queira
	//um desses valores eu posso  colocar um _ no local que seria para uma variavel receber o valor, entao simplesmente nao atribui
	//o resultado retornado a ninguem
	// OBS: EU N USEI := NESTE CASO POIS SO SE PODE ATRIBUIR APOS TER FEITO UMA DECLARAÇÃO EM UMA VARIAVEL, isso é explicado melhor no cmds2.go
	_ , erros = fmt.Println("heelo world", "ola pessoas", "demontrar numeros 100 e ", 100)
	fmt.Println(erros)

}
