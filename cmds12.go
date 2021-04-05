package main

import(
	"fmt"
)
func main(){
	//slice multi-dimensional
  
	ss := [][]int{// declara q dois espacos de slice, o primeiro espco das chaves seria a linha e o segundo espaço a coluna
   // Índice: 0  1  2               // Índice:
		[]int{1, 2, 3},    // 0
		[]int{7, 8, 9},    // 1
		[]int{13, 14,15},  // 2
	}
	 
	
	fmt.Println(ss[2][1])
//                 |   |
//				   |   L->COLUNA
//				   L->LINHA
/*
// ESSE CODIGO É UM EXEMPLO  DE ALGO MAIS COMPLEXO USANDO SLICE
s := [][][][]int{

	[][][]int{
		[][]int{
			[]int{1, 2, 3, 4, 5, 6},
		},
		[][]int{
			[]int{10, 20, 30, 40, 50},
			},
		},
	
	[][][]int{
		[][]int{
			[]int{2, 4, 6, 8, 10},
			},
		[][]int{
			[]int{3},
			},
		},

}
fmt.Println(s[1][0][0][2])
}*/
// ELE RETORNARA O VALOR 6, QUE FOI ESPECIFICADO SUAS CORDENADAS NESSE ULTIMO PRINT	 



//=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=
//maps
// map tem a seguinte estrutura , onde se tem uma identificacao e um valor
//ele é util para performace
fmt.Println("TESTE COM O COMANDO MAPS")

//lado esquerdo é a identificação e do direito o valor atribuido
 lista:= map[int]string{
	1: "joao",
	2: "pedro",
	3: "maria",
 }

fmt.Println(lista)
fmt.Println(lista[1])//quando eu executar isso ele me dara o valor atribuido ao numero colocado entre as chaves
// lembrando que o identificador pode ser uma STRING COMO no ex abaixo
amigos := map[string]int{
	"alfredo": 5551234,
	"joana":   9996674,
}

fmt.Println(amigos)
fmt.Println(amigos["joana"])

// para adicionar algo em alguma desses dois grupos siga esta estrutura

lista[4]= "thiago"
amigos["gopher"] = 444444

fmt.Println("\napos incrementar outro valor em ambas  grupos")

fmt.Println(lista)
fmt.Println(amigos)

//para eu saber  se existe um determinado elemento nesse grupo de elemntos é necessario fazer o seguinte
fmt.Println("\nverificar se existe a identificaçao 6 na lista  e a identificacao jp em amigos\n\n")

//existe,ok := lista[6]

//existe2,ok := lista[6]

//onde esta o _ pode se colocar uma variavel que ira receber o valor da identificação, caso nao exista recebera 0, o ok recebera um valor bool

if _,ok := lista[5] ;false==ok{// é possivel colocar fallse == ok  ou apenas !ok para  dizer se ok é falso, ou eu poderia ter colocado so o ok que diria que ele é true
	fmt.Println("esse elemnto nao esta contido em lista")

}else{
	fmt.Println("esse elemnto esta contido em lista")

}

if _,ok := amigos["dude"] ;!ok{
	fmt.Println("esse elemnto nao esta contido em amigos")

}else{
	fmt.Println("esse elemnto esta contido em amigos")

}

//para deletar use o seguinte o comando

fmt.Println("\nantes de deletar")

fmt.Println(lista)
fmt.Println(amigos)

delete(lista,1)
delete(amigos,"joana")

fmt.Println("\ndepois de deletar")

fmt.Println(lista)
fmt.Println(amigos)
}