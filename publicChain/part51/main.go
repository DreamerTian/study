package main

import (
	"fmt"
	"study/publicChain/part51/BLC"
)

func main()  {
	wallet := BLC.NewWallet()

	address := wallet.GetAddress()

	fmt.Printf("address:%s\n",address)
}
