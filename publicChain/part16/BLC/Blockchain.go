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
	Tip []byte   //最新的区块的Hash
	DB  *bolt.DB //数据库
}

////增加区块到区块链里面
func (blc *Blockchain) AddBlockToBlockchain(data string) {

	err := blc.DB.Update(func(tx *bolt.Tx) error {

		//1.先去获取表
		b := tx.Bucket([]byte(blockTableName))

		if b != nil {
			//先获取最新区块的字节数组
			blockBytes := b.Get(blc.Tip)
			//反序列化
			block := DeserializeBlock(blockBytes)
			//2.创建新区块
			newBlock := NewBlock(data, block.Height + 1, block.Hash)
			//3.将区块序列化 并存储到数据库中
			err := b.Put(newBlock.Hash, newBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}
			//4.更新数据库里面 l 对应的 hash
			err = b.Put([]byte("l"), newBlock.Hash)
			if err != nil {
				log.Panic(err)
			}
			//5.更新blockchain的Tip
			blc.Tip = newBlock.Hash
		}
		return nil
	})

	if err != nil {
		log.Panic(err)
	}
}

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
		//先看看表是否存在
		b := tx.Bucket([]byte(blockTableName))
		//如果表存在
		if b == nil {
			//创建数据表
			b, err = tx.CreateBucket([]byte(blockTableName))
			if err != nil {
				log.Fatal(err)
			}
		}

		if b != nil {
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