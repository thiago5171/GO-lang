package main

import (
	"fmt"
)
var x int
func main(){
/*x=1
for x<=10000{
	fmt.Printf("%v-",x)
	x++
}
*/
//OU
x=1
for {
	fmt.Printf("%v-",x)
	x++
	if x==10000{
		break
	}
}


}