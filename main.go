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
	datas.Init()
	bitquery.Api(config)
	for k, v := range datas.BitcoinDataHashMap {
		fmt.Println("Key:", k, "Value:", v)
	}
	fmt.Println(datas.BitcoinHashList)
}
