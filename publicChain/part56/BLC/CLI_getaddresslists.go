package BLC

import "fmt"

//打印所有的钱包地址
func (cli *CLI)addressLists()  {

	wallets,_ := NewWallets()

	for address,_ := range wallets.WalletsMap{
		fmt.Println("钱包地址:"+address)
	}

}
