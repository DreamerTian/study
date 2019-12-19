package BLC

type TXOutput struct {
	Value int64
	ScriptPubKey string
}

// 判断当前的消费是谁的钱
func (txOutput *TXOutput) UnLockWithAddress(address string) bool{
	return txOutput.ScriptPubKey == address
}