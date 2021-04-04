package main

import (
	"fmt"
)
//'SLICE'
var i int
func main(){
	//diferentemente do array o numero de casas dentro dos colchete é mutavel
	slice:=[]int{1,2,3,4,5,6}
	// ou poderia ser declarado com outro tipo
	slice2:= []string{"maça","laranga","abacaxi"}
	fmt.Println(slice)

	fmt.Println(len(slice))
	fmt.Println(slice[2])
	slice[2]=23221
	fmt.Println(slice[2])
// mas tbm nao é possivel chegar e fazer isso
 //slice[20]=23 pois o slite é feito de arrays e como nao foi criada ainda esse numero de arrays entao n pode requisita-lo

//para acrescenttar um elemento no slice  é mecessario que se use a seguinte estrutura
slice=append(slice, 9,2)// foi acrescentado o 9 e o 2
 //para percorrer o slici é possifavel fazer esses seguintes exemplos, onde o 'valor' recebe o que esta na posição correspondente do slice

	for i, valor := range slice{
		fmt.Printf("no indice %v temos o valor %v \n",i,valor)
	}
	//OU
slice2 = append(slice2, "banana", "uva")
	for _, valor := range slice2{
		
		fmt.Printf(" temos o valor %v  em slice\n",valor)
	 }
	
 // ou é possivel pegar apenas o indices do slice
 for i, _ := range slice{
	fmt.Printf("temos indice %v em slice  \n",i)
}

// para poder recortar uma  parte de um slice vc deve usar a seguine estrutura

// atribua  a uma nova variavel a slice deliimitanto a distancia do corte por [:]
fatia:=slice[:2]
fmt.Println(fatia)
// funciona de forma semelhante ao python o funcionamento nas chaves 



// como excluir um item
fmt.Println(slice)//antes

slice= append(slice[:2],slice[3:]...)//  removendo o valor de indice 2
fmt.Println(slice)//depois da esclusao


//anexando a uma slice
// para anexar um item em uma slice basta colacar essa estrutura
slice = append(slice, 0)// para colocar mais de um item basta colocar a virgula e ir incrementando
fmt.Println(slice)

//para somar uma slice com outra
slice= append(slice, fatia...)
//					        |-> isso siginifica que vai adicionar os elementos  da variavel fatia
//pois como slices sao feitas de arrays entao elas sao de tipos diferente, ou seja, nao conseguem interagir

}