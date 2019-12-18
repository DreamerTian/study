package BLC

type TXInput struct {
	//交易的hash
	TxHash []byte

	Vout int
	//用户名
	ScriptSig string
}
