package datas

import (
	"blockdata-crawling/src/types"
)

var BitcoinBlockDataMap map[string]types.BitcoinBlockData
var BitcoinTxDataMap map[string]types.BitcoinTxData

var BitcoinBlockHashList []string
var BitcoinTxHashList []string

func BitCoinDataStorageInit() {
	BitcoinBlockDataMap = make(map[string]types.BitcoinBlockData)
	BitcoinTxDataMap = make(map[string]types.BitcoinTxData)
}
