package main

import (
	"fmt"
)
//'SLICE'

func main(){
	//diferentemente do array o numero de casas dentro dos colchete é mutavel
	slice:=[]int{1,2,3,4,5,6}
	fmt.Println(slice)

	fmt.Println(len(slice))
	fmt.Println(slice[2])
	slice[2]=23221
	fmt.Println(slice[2])
// mas tbm nao é possivel chegar e fazer isso
 //slice[20]=23 pois o slite é feito de arrays e como nao foi criada ainda esse numero de arrays entao n pode

 
}