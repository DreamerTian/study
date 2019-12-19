package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
	"time"
)

type Block struct {
	//1.区块高度（编号） 第几个区块
	Height int64
	//2.上一个 区块的 HASH
	PrevBlockHash []byte
	//3.交易数据
	Txs []*Transaction
	//4.时间戳
	Timestamp int64
	//5.当前的HASH
	Hash []byte
	//6.Nonce
	Nonce int64
}

// 将区块序列化成字节数组
func (block *Block) Serialize() []byte {
	var result bytes.Buffer

	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(block)

	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

// 需要将txs 转化成字节数组 并返回
// 提供给挖矿是使用 将区块里面所有的交易ID拼接，并且生成hash
func (block *Block) HashTransaction() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _,tx := range block.Txs{
		txHashes = append(txHashes,tx.TxHash)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes,[]byte{}))

	return txHash[:]
}

//反序列化
func DeserializeBlock(blockBytes []byte) *Block {

	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(blockBytes))

	err := decoder.Decode(&block)

	if err != nil {
		log.Panic(err)
	}

	return &block

}

//1.创建新的区块
func NewBlock(txs []*Transaction, height int64, prevBlockHash []byte) *Block {

	//创建区块
	block := &Block{height, prevBlockHash, txs, time.Now().Unix(), nil, 0}
	//调用工作量证明的方法并且返回有效的Hash 和 Nonce
	pow := NewProofOfWork(block)

	//挖矿验证
	hash, nonce := pow.Run()

	block.Hash = hash[:]

	block.Nonce = nonce

	fmt.Println()

	return block
}

//单独写一个方法 生成一个创世区块
func CreateGenesisBlock(txs []*Transaction) *Block {

	//对于一个创世区块来讲 导读可知

	return NewBlock(txs, 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}
