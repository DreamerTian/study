package BLC

import (
	"github.com/boltdb/bolt"
	"log"
)

//迭代器结构体
type BlockchainIterator struct {
	CurrentHash []byte   //当前正在遍历的区块的Hash
	DB          *bolt.DB //数据库
}


func (blockchainIterator *BlockchainIterator) Next() *Block {

	var block *Block

	err := blockchainIterator.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))

		if b != nil {
			currentBlockBytes := b.Get(blockchainIterator.CurrentHash)
			//获取到当前迭代器里面的currentHash所对应的区块
			block = DeserializeBlock(currentBlockBytes)
			//更新迭代器里面的currentHash
			blockchainIterator.CurrentHash = block.PrevBlockHash
		}

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	return block
}
