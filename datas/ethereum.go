package datas

import (
	"blockdata-crawling/src/types"
)

var EthereumBlockDataMap map[string]types.EhtereumData
var EhtereumBlockHashList []string

func EthereumDataStorageInit() {
	EthereumBlockDataMap = make(map[string]types.EhtereumData)
}
