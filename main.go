package main

import (
	"fmt"
	"kyoku-blockchain/wallet"
)

func main() {
	w := wallet.NewWallet()
	fmt.Println(w.PrivateKeyStr())
	// fmt.Println(w.PublicKeyStr())
}
