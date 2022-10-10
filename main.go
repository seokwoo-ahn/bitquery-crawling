package main

import (
	"blockdata-crawling/datas"
	"blockdata-crawling/libs"
	"blockdata-crawling/src/bitquery"
	_ "blockdata-crawling/src/spreadsheet"
	"fmt"
)

func main() {
	config := libs.LoadConfig()
	datas.BitCoinDataStorageInit()
	bitquery.Api(config)
	for k, v := range datas.BitcoinBlockDataMap {
		fmt.Println("Key:", k, "Value:", v)
	}
	fmt.Println(datas.BitcoinBlockHashList)
}
