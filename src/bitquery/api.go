package bitquery

import (
	bitcoin "blockdata-crawling/src/bitquery/bitcoin"
	ethereum "blockdata-crawling/src/bitquery/ethereum"
	"blockdata-crawling/src/types"
)

type scan struct {
	chain    string
	scanType string
}

var scanList []scan

func CallApi(config types.Config) {
	var scan = new(scan)
	scan.chain = config.Chain
	scan.scanType = config.ScanType

	scanList = append(scanList, *scan)

	for _, v := range scanList {
		if v.chain == "bitcoin" {
			bitcoin.BlockScan(config)
		}
		if v.chain == "ethereum" {
			ethereum.BlockScan(config)
		}
	}
}
