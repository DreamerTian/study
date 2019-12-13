package BLC

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	//1.区块高度（编号） 第几个区块
	Height int64
	//2.上一个 区块的 HASH
	PrevBlockHash []byte
	//3.交易数据
	Data []byte
	//4.时间戳
	Timestamp int64
	//5.当前的HASH
	Hash []byte
}

//2.设置当前区块的Hash
func (block *Block) SetHash() {
	// 1.Height []byte
	heightBytes := IntToHex(block.Height)
	// 2. 将时间戳转换为字节数组 []byte
	//  时间戳转换有区别 2代表2进制
	timeString := strconv.FormatInt(block.Timestamp, 2)
	timeBytes := []byte(timeString)
	// 3. 拼接所有属性
	blockBytes := bytes.Join([][]byte{heightBytes, block.PrevBlockHash, block.Data, timeBytes, block.Data}, []byte{})
	// 4.生成Hash
	hash := sha256.Sum256(blockBytes)

	block.Hash = hash[:]
}

//1.创建新的区块
func NewBlock(data string, height int64, prevBlockHash []byte) *Block {

	//创建区块
	block := &Block{height, prevBlockHash, []byte(data), time.Now().Unix(), nil}
	//设置当前区块的hash
	block.SetHash()

	return block
}
