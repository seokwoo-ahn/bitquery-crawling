package bitquery

import (
	"blockdata-crawling/libs/prompts"
	bitcoin "blockdata-crawling/src/bitquery/bitcoin"
	ethereum "blockdata-crawling/src/bitquery/ethereum"
	"blockdata-crawling/src/types"
	"fmt"

	prompt "github.com/c-bata/go-prompt"
)

type scan struct {
	chain    string
	scanType string
}

var scanList []scan

func CallApi(config types.Config) error {
	var scan = new(scan)
	scan.chain = config.Chain
	scanList = append(scanList, *scan)

	for _, v := range scanList {
		fmt.Println(v.chain)
		fmt.Println("Enter Your Scan Type")
		input := prompt.Input("> ", prompts.Completer)

		switch input {
		case "blocks":
			if v.chain == "bitcoin" {
				bitcoin.BlockScan(config)
			} else if v.chain == "ethereum" {
				ethereum.BlockScan(config)
			} else {
				return fmt.Errorf("invalid chain")
			}
		case "transactions":
			fmt.Println("지원 예정")
		case "transaction by hash":
			fmt.Println("지원 예정")
		default:
			return fmt.Errorf("invalid scan type")
		}
	}
	return nil
}
