package mainf

import (
	"fmt"
)

// Em uma constante o TIPO so é definido quando for usado, e pode se adaptar de acordo com oq essa const esta interagindo
const(
	x = 10


// ele retornara  valores de forma sequencial podendo  colocar  _ para anular esse valor mas nao interromper a sequencia
	a = iota  // executando uma operação nessa declaraçao afetara as outras tambem, por exemplo  a= inota +1
	_
	b 
	_
	c  
)

/*OBS: pode se declrar varias constantes utilizando as chaves
const (
	x =10
	y = 45
	z = 66
)
*/
var y float64
var z int


func main(){
	y = x
	fmt.Printf("\n %v %T",y,y)
	z = x
	fmt.Printf("\n%v %T\n",z,z)
// entao o programa retornara 10 sendo inteiro e real de acordo com a situação necessaria	


	fmt.Println(a, b, c)
}