package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
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
func  NewSimpleTransaction(from string, to string, amount int) *Transaction {
	var txInputs []*TXInput
	var txOututs []*TXOutput

	bts ,_ := hex.DecodeString("7ca38cc6335511e454dfa64b748080f2c16f35352a49baa27e2a345ac98641aa")

	//代表消费信息
	txInput := &TXInput{bts,0,from}
	txInputs = append(txInputs,txInput)
	//转账
	txOutput := &TXOutput{int64(amount),to}
	txOututs = append(txOututs,txOutput)
	//找零
	txOutput = &TXOutput{4-int64(amount),from}
	txOututs = append(txOututs,txOutput)

	tx := &Transaction{[]byte{},txInputs,txOututs}

	//设置hash值
	tx.HashTransaction()

	return tx

}

