package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}

func NewBlock(nonce int, previousHash string) *Block {
	return &Block{
		timestamp:    time.Now().UnixNano(),
		nonce:        nonce,
		previousHash: previousHash,
	}
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256([]byte(m))
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64    `json:"timestamp"`
		Nonce        int      `json:"nonce"`
		PreviousHash string   `json:"previoushash"`
		Transactions []string `json:"transactions"`
	}{
		Timestamp:    b.timestamp,
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Transactions: b.transactions,
	})
}

type BlockChain struct {
	transactionPool []string
	chain           []*Block
}

func NewBlockChain() *BlockChain {
	bc := new(BlockChain)
	bc.CreateBlock(0, "init")
	return bc
}

func (bc *BlockChain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

func init() {

}

func main() {
	bc := NewBlockChain()
	bc.CreateBlock(5, "hash 5")
	fmt.Println(bc)
}
