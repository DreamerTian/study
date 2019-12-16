package BLC

import (
	"github.com/boltdb/bolt"
	"fmt"
	"log"
	"math/big"
)

//数据库的名称
const dbName = "blockchain.db"

//表的名称
const blockTableName = "blocks"

type Blockchain struct {
	Tip []byte   //最新的区块的Hash
	DB  *bolt.DB //数据库
}

//遍历输出所有区块的信息
func (blc *Blockchain) Printchain() {
	var block *Block

	var currentHash []byte = blc.Tip
	//1.通过死循环来遍历
	for {

		//查询表
		err := blc.DB.View(func(tx *bolt.Tx) error {

			b := tx.Bucket([]byte(blockTableName))

			if b != nil {
				//拿到最新的区块的字节数组
				blockBytes := b.Get(currentHash)
				//反序列化
				block = DeserializeBlock(blockBytes)
				////1.区块高度（编号） 第几个区块
				//Height int64
				////2.上一个 区块的 HASH
				//PrevBlockHash []byte
				////3.交易数据
				//Data []byte
				////4.时间戳
				//Timestamp int64
				////5.当前的HASH
				//Hash []byte
				////6.Nonce
				//Nonce int64
				fmt.Printf("Height:%d\n",block.Height)
				fmt.Printf("PrevBlockHash:%x\n",block.PrevBlockHash)
				fmt.Printf("Data:%s\n",block.Data)
				fmt.Printf("Timestamp:%d\n",block.Timestamp)
				fmt.Printf("Hash:%x\n",block.Hash)
				fmt.Printf("Nonce:%d\n",block.Nonce)
			}
			return nil
		})

		fmt.Println()

		if err != nil {
			log.Panic(err)
		}

		var hashInt big.Int

		hashInt.SetBytes(block.PrevBlockHash)

		if big.NewInt(0).Cmp(&hashInt) == 0{
			break
		}

		currentHash = block.PrevBlockHash
	}
}

//增加区块到区块链里面
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
		// 返回nil,以便数据库处理相应操作
		return nil
	})
	// 返回区块链对象
	return &Blockchain{blockHash, db}

}
