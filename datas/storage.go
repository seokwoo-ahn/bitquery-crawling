package datas

import (
	"blockdata-crawling/src/types"
)

func StorageInit() {
	BitcoinBlockDataMap = make(map[string]types.BitcoinData)
	EthereumBlockDataMap = make(map[string]types.EhtereumData)
}
