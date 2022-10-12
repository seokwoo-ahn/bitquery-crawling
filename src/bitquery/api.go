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
	chain string
}

var scanList []scan

func CallApi(config types.Config) error {
end:
	for {
		fmt.Println("Enter chain name")
		chain := prompt.Input(">> ", prompts.Chain)
		switch chain {
		case "end":
			break end
		default:
			var scan = new(scan)
			scan.chain = chain
			scanList = append(scanList, *scan)
		}
	}

	for _, v := range scanList {
		fmt.Println(v.chain)
		fmt.Println("Enter Your Scan Type")
		scanType := prompt.Input(">> ", prompts.ScanType)

		switch scanType {
		case "blocks":
			if v.chain == "bitcoin" {
				bitcoin.BlockScan(config)
			} else if v.chain == "ethereum" {
				ethereum.BlockScan(config)
			} else {
				return fmt.Errorf("invalid chain")
			}
		case "transactions":
			if v.chain == "bitcoin" {
				bitcoin.TransactionScan(config)
			} else if v.chain == "ethereum" {
				fmt.Println("지원 예정")
			} else {
				return fmt.Errorf("invalid chain")
			}
		case "transaction by hash":
			fmt.Println("지원 예정")
		default:
			return fmt.Errorf("invalid scan type")
		}
	}
	return nil
}
