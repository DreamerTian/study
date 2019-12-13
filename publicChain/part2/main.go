package main

import (
	"study/publicChain/part2/BLC"
	"fmt"
)

func main(){

	genesisBlock := BLC.CreateGenesisBlock("Genesis Block")

	fmt.Println(genesisBlock)
}
