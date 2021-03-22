package main

import (
	"fmt"
)

var i int
func main(){
// em go existe apenas o FOR como// entaoconheguimos ter uma estrutura equivalente ao whhilw laco de repeticao seguindo uma estrutura parecida com de e apenas adequando as estrutura do for
  for i:=1; i<=10;i++{//repetir de 1 ao 10
    fmt.Printf("%v \n",i)  
    
  }
  
 //para termos algo equivalente ao while em go podemos fazer a seguinte estrutura
i2:=1
for i2<=10 {// entaoconheguimos ter uma estrutura equivalente ao whhile apenas adequando as estrutura do for
  fmt.Printf("%v \n",i2)
  i2++  
}
//e para fazer uma repeticao infinita basta fazer essa estrutura

  // e para sair basta introduzir um laco de repeticao propondo uma condicao e caso nao atenda pode usar a funcao break que sai da estrutura for
  // como para quebrar o loop todo se usa break, quando queremos   quebrar apenas uma iteracao especifica do loop usammos  o 'continue'

}
