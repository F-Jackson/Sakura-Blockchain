package main

import (
	"io"
	"kyoku-blockchain/block"
	"kyoku-blockchain/wallet"
	"log"
	"net/http"
	"strconv"
)

var cache map[string]*block.BlockChain = make(map[string]*block.BlockChain)

type BlockChainServer struct {
	port uint16
}

func NewBlockChainServer(port uint16) *BlockChainServer {
	return &BlockChainServer{port}
}

func (bcs *BlockChainServer) Port() uint16 {
	return bcs.port
}

func (bcs *BlockChainServer) GetBlockChain() *block.BlockChain {
	bc, ok := cache["blockchain"]
	if !ok {
		minersWallet := wallet.NewWallet()
		bc = block.NewBlockChain(
			minersWallet.BlockChainAddress(),
			bcs.Port())
		cache["blockchain"] = bc
	}
	return bc
}

func HelloWorld(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello world")
}

func (bcs *BlockChainServer) GetChain(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		w.Header().Add("Content-Type", "application/json")
		bc := bcs.GetBlockChain()
		m, _ := bc.MarshalJSON()
		io.WriteString(w, string(m[:]))
	default:
		log.Panicln("Error needs to be get method")
	}
}

func (bcs *BlockChainServer) Run() {
	http.HandleFunc("/", bcs.GetChain)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(bcs.Port())), nil))
}
