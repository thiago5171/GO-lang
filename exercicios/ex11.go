package main

import (
	"fmt"
)

func main(){
	resto:=0
	for x:=10;x<=100;x++{
		resto=x%4
		fmt.Printf("%v-%v\n",x,resto)
	}

}
