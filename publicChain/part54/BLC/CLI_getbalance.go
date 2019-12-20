package BLC

import "fmt"

//用它去查询余额
func (cli *CLI) getBalance(address string) {

	blockchain := BlockchainObject()

	defer blockchain.DB.Close()

	amount := blockchain.GetBalance(address)

	fmt.Println(amount)
}

