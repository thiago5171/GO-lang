package main
// tente acessar todos os itens de uma slice *sem* utilizar range.
import (
	"fmt"
)

func main(){

	s :=[]int{1,2,3,4,5}

	for i:=0;i< len(s);i++{
 		fmt.Println(s[i])
	} 
}
