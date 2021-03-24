package main

import (
	"fmt"
)

func main(){
	/*o println tem o objetivo de  exibir e pular uma linha
	ja o printf nao da um enter para ir para proxima linha, e nela podemos usar como se fosse as chaves{} e o .format de python
	aqui usamos um desses porcentagengens e uma letra especificada junto para cada situação
	como por exemplo:
	e onde for colocado o %v é exatamente onde ira ficar a variavel
	fmt.ptintf(" a idade dele é %v",idade)

	%v : para valor
	%T : para tipo da variável
	%d : para números decimais
	%b : para números binários 
	%#u : para unicode 
	%#x : para números hexadecimal 
	*/
	idade := 19
	nome := "thiago"
	fmt.Printf("%v tem %v anos de idade \n",nome, idade)
	// lembrando que se a variavel nao estiver sendo usada o programa nao vai rodar
	//----------------------------------------------------------------------------------------------------
	// em GO quando atribuimos um tipo primitivo em uma variavel não é possivel trocar esse tipo posteriormente
	//----------------------------------------------------------------------------------------------------
	//so se pode fazer uma atribuição para tal variavel sem falar o tipo primitivo (ou seja uma declaração) 
	//dentro de uma  func main(){}
	//----------------------------------------------------------------------------------------------------
	//usando o /t dentro das aspas teremos um TAB nesse espaço
	//para aparecer as aspas no texo tem que usar contra barra antes \" texto desejado \"
	//----------------------------------------------------------------------------------------------------
	//caso queira que mostre um texto CRU renm nenhuma interpretação , basta colocar set texto dentro de acentos grave
	// `texto sem formatação `

	
	
}
