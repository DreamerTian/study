package BLC

import (
	"fmt"
	"os"
)

func (cli *CLI) addBlock(txs []*Transaction) {

	if DbExists() == false {
		fmt.Println("请先创建创世区块.....")
		os.Exit(1)
	}

	blockchain := BlockchainObject()
	defer blockchain.DB.Close()

	blockchain.AddBlockToBlockchain(txs)
}
