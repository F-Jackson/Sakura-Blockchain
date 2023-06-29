package blockchain_server

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

type BlockChainServer struct {
	port uint16
}

func NewBlockChainServer(port uint16) *BlockChainServer {
	return &BlockChainServer{port}
}

func (bcs *BlockChainServer) Port() uint16 {
	return bcs.port
}

func HelloWorld(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello world")
}

func (bcs *BlockChainServer) Run() {
	http.HandleFunc("/", HelloWorld)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(bcs.Port())), nil))
}
