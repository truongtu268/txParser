package main

import (
	"encoding/json"
	"fmt"
)

type Parser interface {
	// last parsed block
	GetCurrentBlock() int
	// add address to observer
	Subscribe(address string) bool
	// list of inbound or outbound transactions for an address
	GetTransactions(address string) []Transaction
}

type StorageTrans struct {
	trans        []Transaction
	mapTransHash map[string]string
}

type TxParser struct {
	currentBlock      int64
	currentBlockCount int
	currentHash       string
	mapAdrAndTrans    map[string]StorageTrans
}

func NewTxParser() *TxParser {
	return &TxParser{mapAdrAndTrans: make(map[string]StorageTrans)}
}

func (parser *TxParser) GetCurrentBlock() int64 {
	resByte, _ := callToEth(GetCurrentBlock, defaultParams)
	res := &GetCurrentBlockResponse{}
	_ = json.Unmarshal(resByte, &res)
	currentBlock, _ := convertHexToInt(res.Result)
	return currentBlock
}

func (parser *TxParser) Subscribe(address string) bool {
	if _, ok := parser.mapAdrAndTrans[address]; ok {
		return false
	}
	parser.mapAdrAndTrans[address] = StorageTrans{
		mapTransHash: make(map[string]string),
		trans:        make([]Transaction, 0),
	}
	return true
}

func (parser *TxParser) collectTransactions() {
	blockNo := parser.GetCurrentBlock()

	preHash, sourceTrans, _ := parser.getTransactionByNumber(blockNo)
	parser.filterTransByAddress(sourceTrans)
	if preHash == "" {
		return
	}
	if blockNo == parser.currentBlock {
		return
	}
	for preHash != "" {
		preHash, sourceTrans, _ = parser.getTransactionByHash(preHash)
		parser.filterTransByAddress(sourceTrans)
	}
	return
}

func (parser *TxParser) getTransactionByNumber(blockNo int64) (previousHash string, trans []Transaction, err error) {
	params := make([]any, 0)
	hex := convertIntToHex(blockNo)
	params = append(params, hex, true)
	resByte, _ := callToEth(GetBlockInfoByNumber, params)
	res := &GetBlockResponse{}
	_ = json.Unmarshal(resByte, &res)
	previousHash = res.Result.ParentHash
	trans = res.Result.Transactions
	return
}

func (parser *TxParser) getTransactionByHash(hashStr string) (previousHash string, trans []Transaction, err error) {
	params := make([]any, 0)
	params = append(params, hashStr, true)
	resByte, _ := callToEth(GetBlockInfoByHash, params)
	res := &GetBlockResponse{}
	_ = json.Unmarshal(resByte, &res)
	previousHash = res.Result.ParentHash
	trans = res.Result.Transactions
	return
}

func (parser *TxParser) filterTransByAddress(source []Transaction) {
	for _, transaction := range source {
		fromStorage, fromOk := parser.mapAdrAndTrans[transaction.From]
		toStorage, toOk := parser.mapAdrAndTrans[transaction.To]
		if !fromOk && !toOk {
			continue
		}
		if fromOk {
			_, inMap := fromStorage.mapTransHash[transaction.Hash]
			if !inMap {
				fromStorage.mapTransHash[transaction.Hash] = transaction.Hash
				fromStorage.trans = append(fromStorage.trans, transaction)
				parser.mapAdrAndTrans[transaction.From] = fromStorage
			}
		}
		if toOk {
			_, inMap := toStorage.mapTransHash[transaction.Hash]
			if !inMap {
				toStorage.mapTransHash[transaction.Hash] = transaction.Hash
				toStorage.trans = append(toStorage.trans, transaction)
				parser.mapAdrAndTrans[transaction.To] = toStorage
			}
		}
	}
}

func (parser *TxParser) GetTransactions(address string) []Transaction {
	storage, ok := parser.mapAdrAndTrans[address]
	if !ok {
		return nil
	}
	return storage.trans
}

func main() {
	txParser := NewTxParser()
	txParser.Subscribe("0x3fc91a3afd70395cd496c647d5a6cc9d4b2b7fad")
	txParser.collectTransactions()
	trans := txParser.GetTransactions("0x3fc91a3afd70395cd496c647d5a6cc9d4b2b7fad")
	byts, _ := json.Marshal(trans)
	fmt.Println(string(byts))
}
