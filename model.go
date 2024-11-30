package main

type Transaction struct {
	Type                 string   `json:"type"`
	ChainID              string   `json:"chainId"`
	Nonce                string   `json:"nonce"`
	Gas                  string   `json:"gas"`
	MaxFeePerGas         string   `json:"maxFeePerGas,omitempty"`
	MaxPriorityFeePerGas string   `json:"maxPriorityFeePerGas,omitempty"`
	To                   string   `json:"to"`
	Value                string   `json:"value"`
	AccessList           []any    `json:"accessList,omitempty"`
	Input                string   `json:"input"`
	R                    string   `json:"r"`
	S                    string   `json:"s"`
	YParity              string   `json:"yParity,omitempty"`
	V                    string   `json:"v"`
	Hash                 string   `json:"hash"`
	BlockHash            string   `json:"blockHash"`
	BlockNumber          string   `json:"blockNumber"`
	TransactionIndex     string   `json:"transactionIndex"`
	From                 string   `json:"from"`
	GasPrice             string   `json:"gasPrice"`
	BlobVersionedHashes  []string `json:"blobVersionedHashes,omitempty"`
	MaxFeePerBlobGas     string   `json:"maxFeePerBlobGas,omitempty"`
}

type Withdrawal struct {
	Index          string `json:"index"`
	ValidatorIndex string `json:"validatorIndex"`
	Address        string `json:"address"`
	Amount         string `json:"amount"`
}

type Block struct {
	Hash                  string        `json:"hash"`
	ParentHash            string        `json:"parentHash"`
	Sha3Uncles            string        `json:"sha3Uncles"`
	Miner                 string        `json:"miner"`
	StateRoot             string        `json:"stateRoot"`
	TransactionsRoot      string        `json:"transactionsRoot"`
	ReceiptsRoot          string        `json:"receiptsRoot"`
	LogsBloom             string        `json:"logsBloom"`
	Difficulty            string        `json:"difficulty"`
	Number                string        `json:"number"`
	GasLimit              string        `json:"gasLimit"`
	GasUsed               string        `json:"gasUsed"`
	Timestamp             string        `json:"timestamp"`
	ExtraData             string        `json:"extraData"`
	MixHash               string        `json:"mixHash"`
	Nonce                 string        `json:"nonce"`
	BaseFeePerGas         string        `json:"baseFeePerGas"`
	WithdrawalsRoot       string        `json:"withdrawalsRoot"`
	BlobGasUsed           string        `json:"blobGasUsed"`
	ExcessBlobGas         string        `json:"excessBlobGas"`
	ParentBeaconBlockRoot string        `json:"parentBeaconBlockRoot"`
	TotalDifficulty       string        `json:"totalDifficulty"`
	Size                  string        `json:"size"`
	Uncles                []any         `json:"uncles"`
	Transactions          []Transaction `json:"transactions"`
	Withdrawals           []Withdrawal  `json:"withdrawals"`
}

type GetBlockResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  Block  `json:"result"`
}

type Request struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  []any  `json:"params"`
	ID      int    `json:"id"`
}

type GetCurrentBlockResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  string `json:"result"`
	ID      int    `json:"id"`
}
