package BLC

type Blockchain struct {
	//存储有序的区块
	Blocks []*Block
}

//1.创建带有创世区块的区块链
func CreateBlockchainWithGenesisBlock() *Blockchain {

	//创建创世区块
	genesisBlock := CreateGenesisBlock("Genesis Data.......")

	// 返回区块链对象
	return &Blockchain{[]*Block{genesisBlock}}
}
