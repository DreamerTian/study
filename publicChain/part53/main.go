package main

import (
	"fmt"
	"study/publicChain/part53/BLC"
)

func main()  {

	wallets := BLC.NewWallets()

	fmt.Println(wallets.Wallets)

	wallets.CreateNewWallet()

}
