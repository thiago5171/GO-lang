package main

import (
	"fmt"
)

func main(){

// switch
/* ele seria um if , porem diferente  estruturalmente, onde se coloca 'case' e  a setença e oque sera executado

ele analisa de cima para baixo buscando uma opção correta, e quando ele encontra a primeira correta  ele executa e nao continua
buscamdo outros casos
*/
fmt.Println("\n1°exemplo\n")
x:=100
switch{
case x==100:
	fmt.Println("x=100")
case x<=101:
	fmt.Println("x<101")
case x >= 99:
	fmt.Println("x>99")	
}

fmt.Println("\n2°exemplo\n")

//USO COM  STRINGS
t:= "p"
switch t {
case "p":
	fmt.Println("p=t")
case "t":
fmt.Println("t=t")
case "r":
	fmt.Println("r=t")
}

fmt.Println("\n3°exemplo\n")

// 'fallthrough'--> ao usar esse comando no switch em um caso, e esse caso for selecionado ele vai executar a 
// declaração do caso escolhido do caso abaixo logo em seguida sem fazer uma analise (não podendo ser o ultimo caso)

switch {
case x==100:
	fmt.Println("x=100")
case x <= 101:
	fmt.Println("x<101")
	fallthrough
case x >= 99:
	fmt.Println("x>99")	
	}

fmt.Println("\n4°exemplo\n")
//'default' esse comando é usado no final do switch para que quando nao seja acionado nenhum dos casos propopostos
// entao ele ira executar o default como ultima opção
x2:=1000
switch  {
case x2==100:
	fmt.Println("x2 =100")
case x2 <99:
	fmt.Println("x2 <99")	
default:
	fmt.Println("x2 não é nenhum dos casos precolocados")	
}

fmt.Println("\n5°exemplo\n")
//'cases compostos' é possivel colocar em um case uma ou mais declarações , lembrando que sera tratado ele ira usar o 'ou' para verificar as declarações dentro do case, como por exemplo
escalacao:= "jose"
switch escalacao{
case "jose", "maria":
	fmt.Println("hoje quem esta na firma é o time 1")
case "pedro", "joao":
	fmt.Println("hoje quem esta na firma é o time 2")
 }

fmt.Println("\nou pode ser usado tbm desse jeito:\n")
i:=18
switch{
case (i==1), (i==2):
	fmt.Println("1 ou 2")
case (i==3), (i==4):
	fmt.Println("3 ou 4")
case (x>10):
	fmt.Println("maior q 10")

	}
} 

