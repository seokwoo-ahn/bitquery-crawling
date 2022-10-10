package datas

import (
	"blockdata-crawling/src/types"
)

var BitcoinBlockDataMap map[string]types.BitcoinData
var BitcoinBlockHashList []string

func BitCoinDataStorageInit() {
	BitcoinBlockDataMap = make(map[string]types.BitcoinData)
}
