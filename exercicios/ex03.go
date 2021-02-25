package main

import (
	"fmt"
)

var y = "james bond"
var z = true

func main(){
	
	s := fmt.Sprintf(y,z)
	fmt.Println(s)

}