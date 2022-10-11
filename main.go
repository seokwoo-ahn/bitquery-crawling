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
	datas.StorageInit()

	if err := bitquery.CallApi(config); err != nil {
		fmt.Println(err)
		return
	}

	for k, v := range datas.BitcoinBlockDataMap {
		fmt.Println("Key:", k, "Value:", v)
	}
	for k, v := range datas.EthereumBlockDataMap {
		fmt.Println("Key:", k, "Value:", v)
	}
	fmt.Println(datas.BitcoinBlockHashList)
	fmt.Println(datas.EhtereumBlockHashList)
}
