package main

import (
	"study/publicChain/part3/BLC"
	"fmt"
)

func main(){

	genesisBlickchain := BLC.CreateBlockchainWithGenesisBlock()

	fmt.Println(genesisBlickchain)

	fmt.Println(genesisBlickchain.Blocks)
	fmt.Println(genesisBlickchain.Blocks[0])
}
