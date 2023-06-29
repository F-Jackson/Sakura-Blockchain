package main

import (
	"encoding/json"
	"io"
	"kyoku-blockchain/wallet"
	"log"
	"net/http"
	"path"
	"strconv"
	"text/template"
)

const TEMP_DIR = "wallet_server/templates"

type WalletServer struct {
	port    uint16
	gateway string
}

func NewWalletServer(port uint16, gateway string) *WalletServer {
	return &WalletServer{port, gateway}
}

func (ws *WalletServer) Port() uint16 {
	return ws.port
}

func (ws *WalletServer) Gateway() string {
	return ws.gateway
}

func (ws *WalletServer) Index(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		t, _ := template.ParseFiles(path.Join(TEMP_DIR, "index.html"))
		t.Execute(w, "")
	default:
		log.Printf("Error invalid http method")
	}
}

func (ws *WalletServer) Wallet(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		decoder := json.NewDecoder(req.Body)
		var t wallet.TransactionRequest
		err := decoder.Decode(&t)
		if err != nil {
			log.Printf("Error %v", err)
			io.WriteString(w, "ERROR")
			return
		}
		if !t.Validate() {
			log.Println("Error missing fields")
			io.WriteString(w, "ERROR")
			return
		}

	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error invalid http method")
	}
}

func (ws *WalletServer) Run() {
	http.HandleFunc("/", ws.Index)
	http.HandleFunc("/wallet", ws.Wallet)
	log.Fatal(http.ListenAndServe("0.0.0.0"+strconv.Itoa(int(ws.Port())), nil))
}
