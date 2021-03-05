package main

import (
	"fmt"
)
var a,b,c,d,e,f bool
func main(){

	a = 5 == 5
	b = 5 != 6
	c = 5 <= 7
	d = 5 < 8
	e = 5 >= 6
	f = 5 > 4
	fmt.Println(a,b,c,d,e,f)
}