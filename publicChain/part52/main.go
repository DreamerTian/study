package main

import (
	"fmt"
	"study/publicChain/part52/BLC"
)

func main()  {
	wallet := BLC.NewWallet()

	address := wallet.GetAddress()

	isValid := BLC.IsValidForAddress(address)

	fmt.Printf("address:%s\n",address)
	fmt.Printf("address:%s,为：%v\n",address,isValid)
}
