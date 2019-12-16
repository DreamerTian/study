package main

import (
	"study/publicChain/part17/BLC"
)

func main() {
	//创世区块
	blockchain := BLC.CreateBlockchainWithGenesisBlock()
	defer blockchain.DB.Close()
	//新区块
	blockchain.AddBlockToBlockchain("Send 100RMB to zhangqiang1")
	blockchain.AddBlockToBlockchain("Send 120RMB to zhangqiang2")
	blockchain.AddBlockToBlockchain("Send 130RMB to zhangqiang3")
	blockchain.AddBlockToBlockchain("Send 140RMB to zhangqiang4")

	blockchain.Printchain()
}
