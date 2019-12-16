package main

import (
	"study/publicChain/part15/BLC"
)

func main() {
	//创世区块
	blockchain := BLC.CreateBlockchainWithGenesisBlock()
	defer blockchain.DB.Close()
	//新区块
	//blockchain.AddBlockToBlockchain("Send 100RMB to zhangqiang1", blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	//blockchain.AddBlockToBlockchain("Send 120RMB to zhangqiang2", blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	//blockchain.AddBlockToBlockchain("Send 130RMB to zhangqiang3", blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	//blockchain.AddBlockToBlockchain("Send 140RMB to zhangqiang4", blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)

	//fmt.Println(blockchain)
}
