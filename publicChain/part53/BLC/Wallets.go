package BLC

import "fmt"

type Wallets struct {
	Wallets map[string]*Wallet
}

//创建钱包的集合
func NewWallets() *Wallets{

	wallets := &Wallets{}

	wallets.Wallets = make(map[string]*Wallet)

	return wallets
}

func (w *Wallets) CreateNewWallet()  {
	wallet := NewWallet()
	fmt.Printf("%s\n",wallet.GetAddress())

	w.Wallets[string(wallet.GetAddress())] = wallet
}
