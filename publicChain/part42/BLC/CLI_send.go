package BLC

import "os"

//转账
func (cli *CLI) send(from []string, to []string, amount []string) {

	if DbExists() == false{
		os.Exit(1)
	}

	blockchain := BlockchainObject()

	defer blockchain.DB.Close()

	blockchain.MineNewBlock(from,to,amount)

}
