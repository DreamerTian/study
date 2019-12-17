package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

// UTXO
type Transaction struct {
	//1.交易hash
	TxHash []byte
	//2.输入
	Vins []*TXInput
	//3.输出
	Vouts []*TXOutput
}

// Transaction 创建分两种情况
//1. 创世区块创建的时候  （特殊）
func NewCoinbaseTransaction(address string) *Transaction{

	//代表消费信息
	txInput := &TXInput{[]byte{},-1,"Genesis Data"}
	//
	txOutput := &TXOutput{10,address}

	txCoinbase := &Transaction{[]byte{},[]*TXInput{txInput},[]*TXOutput{txOutput}}

	//设置hash值
	txCoinbase.HashTransaction()

	return txCoinbase
}

// 将区块序列化成字节数组
func (tx *Transaction) HashTransaction() {

	var result bytes.Buffer

	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(tx)

	if err != nil {
		log.Panic(err)
	}

	hash := sha256.Sum256(result.Bytes())

	tx.TxHash = hash[:]
}

//2. 转账的时候产生

