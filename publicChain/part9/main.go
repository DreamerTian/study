package main

import (
	"fmt"
	"study/publicChain/part9/BLC"
)

func main() {

	block := BLC.NewBlock("Test",1,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})

	fmt.Printf("%d\n",block.Nonce)
	fmt.Printf("%x\n",block.Hash)

	pow := BLC.NewProofOfWork(block)

	fmt.Printf("%v",pow.IsValid())
}
