package main

import (
	"blockdata-crawling/libs"
	"blockdata-crawling/src/bitquery"
	_ "blockdata-crawling/src/spread_sheet"
)

func main() {
	dataSource := libs.LoadDataSourceConfig()
	bitquery.Api(dataSource)
}
