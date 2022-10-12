package datas

import (
	"blockdata-crawling/src/types"
)

var EthereumBlockDataMap map[string]types.EthereumBlockData
var EthereumBlockHashList []string

func EthereumDataStorageInit() {
	EthereumBlockDataMap = make(map[string]types.EthereumBlockData)
}
