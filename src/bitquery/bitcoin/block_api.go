package bitcoin

import (
	"blockdata-crawling/datas"
	"blockdata-crawling/src/types"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func BlockScan(config types.Config) {
	url := config.Datasource.Url
	method := "POST"
	jsonData := map[string]string{
		"query": `
                query MyQuery{
            		bitcoin(network: bitcoin){
                		blocks(date: {after: "2022-09-19"}){
							timestamp {
								time
								unixtime
							}
							height
							blockHash
							difficulty
							transactionCount
						}
                	}
                }
        `,
		"variables": "{}",
	}
	jsonValue, _ := json.Marshal(jsonData)
	payload := bytes.NewBuffer(jsonValue)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-API-KEY", config.Datasource.ApiKey)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		panic("bad request!")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	resultObj := make(map[string]interface{})
	json.Unmarshal(body, &resultObj)

	data := resultObj["data"].(map[string]interface{})
	bitcoin := data["bitcoin"].(map[string]interface{})
	blocks := bitcoin["blocks"].([]interface{})
	for _, v := range blocks {
		var blockData types.BitcoinData
		var timeStamp types.Timestamp
		var blockHash string
		dataMap := v.(map[string]interface{})
		blockHash = dataMap["blockHash"].(string)
		timeStampMap := dataMap["timestamp"].(map[string]interface{})
		timeStamp.Time = timeStampMap["time"].(string)
		timeStamp.Unixtime = timeStampMap["unixtime"].(float64)

		blockData.Timestamp = timeStamp
		blockData.Height = dataMap["height"].(float64)
		blockData.BlockHash = blockHash
		blockData.Difficulty = dataMap["difficulty"].(float64)
		blockData.TransactionCount = dataMap["transactionCount"].(float64)

		datas.BitcoinBlockDataMap[blockHash] = blockData
		datas.BitcoinBlockHashList = append(datas.BitcoinBlockHashList, blockHash)
	}
}
