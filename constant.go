package main

type MethodName string

const (
	GetCurrentBlock      MethodName = "eth_blockNumber"
	GetBlockInfoByNumber MethodName = "eth_getBlockByNumber"
	GetBlockInfoByHash   MethodName = "eth_getBlockByHash"

	ethereumUrl = "https://ethereum-rpc.publicnode.com"
)

var defaultParams []any
