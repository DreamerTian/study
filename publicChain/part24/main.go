package main

import (
	"study/publicChain/part24/BLC"
)

func main() {
	//创世区块
	blockchain := BLC.CreateBlockchainWithGenesisBlock()
	defer blockchain.DB.Close()

	cli := &BLC.CLI{blockchain}

	cli.Run()
}
