package BLC

type TXInput struct {
	//交易的hash
	TxHash []byte

	Vout int
	//用户名
	ScriptSig string
}

// 判断当前的消费是谁的钱
func (txInput *TXInput) UnLockWithAddress(address string) bool{
	return txInput.ScriptSig == address
}
