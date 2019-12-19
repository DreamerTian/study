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

//判断导航全的交易是否是Coinbase交易
func (tx *Transaction)IsCoinbaseTransaction() bool  {
	return len(tx.Vins[0].TxHash) == 0 && tx.Vins[0].Vout == -1
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
func  NewSimpleTransaction(from string, to string, amount int,blockchain *Blockchain,txs []*Transaction) *Transaction {

	//2.通过一个函数返回两个值
	money,spendableUTXODic := blockchain.FindSpendableUTXOS(from,amount,txs)

	var txInputs []*TXInput
	var txOututs []*TXOutput

	for txHash,indexArray := range spendableUTXODic{
		txHashBytes,_ := hex.DecodeString(txHash)
		for _,index := range indexArray {
			//代表消费信息
			txInput := &TXInput{txHashBytes,index,from}
			txInputs = append(txInputs,txInput)
		}
	}
	//转账
	txOutput := &TXOutput{int64(amount),to}
	txOututs = append(txOututs,txOutput)
	//找零
	txOutput = &TXOutput{int64(money)-int64(amount),from}
	txOututs = append(txOututs,txOutput)

	tx := &Transaction{[]byte{},txInputs,txOututs}

	//设置hash值
	tx.HashTransaction()

	return tx

}

