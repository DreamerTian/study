package BLC

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"math/big"
	"os"
	"strconv"
	"time"
)

//数据库的名称
const dbName = "blockchain.db"

//表的名称
const blockTableName = "blocks"

type Blockchain struct {
	Tip []byte   //最新的区块的Hash
	DB  *bolt.DB //数据库
}

//迭代器方法
func (blc *Blockchain) Iterator() *BlockchainIterator {
	return &BlockchainIterator{blc.Tip, blc.DB}
}

//遍历输出所有区块的信息
func (blc *Blockchain) Printchain() {

	blockchainIterator := blc.Iterator()

	for {
		block := blockchainIterator.Next()

		fmt.Printf("Height:%d\n", block.Height)
		fmt.Printf("PrevBlockHash:%x\n", block.PrevBlockHash)
		fmt.Printf("Timestamp:%s\n", time.Unix(block.Timestamp, 0).Format("2006-01-02 03:04:05 PM"))
		fmt.Printf("Hash:%x\n", block.Hash)
		fmt.Printf("Nonce:%d\n", block.Nonce)
		fmt.Println("Txs")
		for _, tx := range block.Txs {
			fmt.Printf("%x\n", tx.TxHash)
			fmt.Println("Vins:")
			for _, in := range tx.Vins {
				fmt.Printf("%x\n", in.TxHash)
				fmt.Printf("%d\n", in.Vout)
				fmt.Printf("%s\n", in.ScriptSig)
			}
			fmt.Println("Vouts:")
			for _, out := range tx.Vouts {
				fmt.Printf("%d\n", out.Value)
				fmt.Printf("%s\n", out.ScriptPubKey)
			}
		}

		fmt.Println("-----------------------------------------------------")

		var hashInt big.Int

		hashInt.SetBytes(block.PrevBlockHash)

		if big.NewInt(0).Cmp(&hashInt) == 0 {
			break
		}
	}
}

//增加区块到区块链里面
func (blc *Blockchain) AddBlockToBlockchain(txs []*Transaction) {

	err := blc.DB.Update(func(tx *bolt.Tx) error {

		//1.先去获取表
		b := tx.Bucket([]byte(blockTableName))

		if b != nil {
			//先获取最新区块的字节数组
			blockBytes := b.Get(blc.Tip)
			//反序列化
			block := DeserializeBlock(blockBytes)
			//2.创建新区块
			newBlock := NewBlock(txs, block.Height+1, block.Hash)
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
func CreateBlockchainWithGenesisBlock(address string) *Blockchain {

	if DbExists() {
		fmt.Println("创世区块已经存在......")
		os.Exit(1)
	}
	fmt.Println("正在创建创世区块......")
	//创建或者打开数据库
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	var genesisHash []byte
	//创建表
	err = db.Update(func(tx *bolt.Tx) error {

		//创建数据表
		b, err := tx.CreateBucket([]byte(blockTableName))
		if err != nil {
			log.Fatal(err)
		}

		if b != nil {
			//创建创世区块

			//创建一个conibase
			txCoinbase := NewCoinbaseTransaction(address)

			genesisBlock := CreateGenesisBlock([]*Transaction{txCoinbase})

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
			genesisHash = genesisBlock.Hash
		}
		// 返回nil,以便数据库处理相应操作
		return nil
	})

	return &Blockchain{genesisHash, db}
}

//判断数据库是否已经存在
func DbExists() bool {
	if _, err := os.Stat(dbName); os.IsNotExist(err) {
		return false
	}
	return true
}

//返回Blockchain 对象
func BlockchainObject() *Blockchain {

	//创建或者打开数据库
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	var tip []byte
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b != nil {
			tip = b.Get([]byte("l"))
		}
		return nil
	})

	return &Blockchain{tip, db}
}

//转账并形成新的区块
func (blc *Blockchain) MineNewBlock(from []string, to []string, amount []string) {

	//1.建立一笔交易

	value ,_ :=strconv.Atoi(amount[0])

	tx := NewSimpleTransaction(from[0],to[0],value)

	fmt.Println(from)
	fmt.Println(to)
	fmt.Println(amount)

	//1.通过相关算法建立交易数组  Transaction数组

	var txs []*Transaction
	txs = append(txs, tx)

	var block *Block

	err := blc.DB.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(blockTableName))

		if b != nil {
			hash := b.Get([]byte("l"))

			blockBytes := b.Get(hash)

			block = DeserializeBlock(blockBytes)
		}

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	//2.建立新的区块
	newBlock := NewBlock(txs, block.Height+1, block.Hash)

	//将新区快存储到数据库
	err = blc.DB.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(blockTableName))

		if b != nil {
			b.Put(newBlock.Hash,newBlock.Serialize())
			b.Put([]byte("l"),newBlock.Hash)

			blc.Tip = newBlock.Hash
		}

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

}
