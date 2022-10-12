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

func TransactionScan(config types.Config) {
	url := config.Datasource.Url
	method := "POST"
	jsonData := map[string]string{
		"query": `
                query MyQuery{
            		bitcoin(network: bitcoin){
                		transactions{
							block(height: {gt: 758340}){
								timestamp {
									time
									unixtime
								}
								height								
							}
							count
							hash
							inputValue
							outputValue
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
	transactions := bitcoin["transactions"].([]interface{})
	for _, v := range transactions {
		var txData types.BitcoinTxData
		var blockData types.BitcoinBlockData
		var timeStamp types.Timestamp

		dataMap := v.(map[string]interface{})

		blockMap := dataMap["block"].(map[string]interface{})
		timeStampMap := blockMap["timestamp"].(map[string]interface{})
		timeStamp.Time = timeStampMap["time"].(string)
		timeStamp.Unixtime = timeStampMap["unixtime"].(float64)
		height := blockMap["height"].(float64)
		count := dataMap["count"].(float64)
		hash := dataMap["hash"].(string)
		inputValue := dataMap["inputValue"].(float64)
		outputValue := dataMap["outputValue"].(float64)

		blockData.Timestamp = timeStamp
		blockData.Height = height

		txData.BitcoinBlockData = blockData
		txData.Count = count
		txData.Hash = hash
		txData.InputValue = inputValue
		txData.OutputValue = outputValue

		datas.BitcoinTxDataMap[hash] = txData
		datas.BitcoinTxHashList = append(datas.BitcoinTxHashList, hash)
	}
}
