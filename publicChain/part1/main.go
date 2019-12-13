package main

import (
	"study/publicChain/part1/BLC"
	"fmt"
)

func main(){
	block := BLC.NewBlock("Genenis Block", 1,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})

	fmt.Println(block)
}
