package datas

import (
	"blockdata-crawling/src/types"
)

var BitcoinDataHashMap map[string]types.BitcoinData
var BitcoinHashList []string

func Init() {
	BitcoinDataHashMap = make(map[string]types.BitcoinData)
}
