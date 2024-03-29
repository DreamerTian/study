package BLC

import (
	"github.com/boltdb/bolt"
	"log"
)

//数据库的名称
const dbName = "blockchain.db"

//表的名称
const blockTableName = "blocks"

type Blockchain struct {
	Tip []byte  //最新的区块的Hash
	DB  *bolt.DB //数据库
}

////增加区块到区块链里面
//func (blc *Blockchain) AddBlockToBlockchain(data string, height int64, preHash []byte) {
//	//创建新区块
//	newBlock := NewBlock(data, height, preHash)
//	// 往链里面添加新的区块
//	blc.Blocks = append(blc.Blocks, newBlock)
//}

//1.创建带有创世区块的区块链
func CreateBlockchainWithGenesisBlock() *Blockchain {

	//创建或者打开数据库
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	var blockHash []byte

	//创建表
	err = db.Update(func(tx *bolt.Tx) error {
		//创建数据表
		b, err := tx.CreateBucket([]byte(blockTableName))
		if err != nil {
			log.Fatal(err)
		}

		if b == nil {
			//创建创世区块
			genesisBlock := CreateGenesisBlock("Genesis Data.......")
			//将创世区块序列化后存储到表当中
			err := b.Put(genesisBlock.Hash, genesisBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}
			//存储最新的区块的hash
			err = b.Put([]byte("l"), genesisBlock.Hash)
			if err != nil {
				log.Panic(err)
			}

			blockHash = genesisBlock.Hash
		}
		// 返回nil,以便数据库处理相应操作
		return nil
	})
	// 返回区块链对象
	return &Blockchain{blockHash, db}


}
