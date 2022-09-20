package graphql

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
)

type Data struct {
	Data []Chain
}

type Chain struct {
	Chain []Result
}

type Result struct {
	Network       string
	Timestamptime string
	Blocksheight  int
}

func Api() {
	url := "https://graphql.bitquery.io"
	method := "POST"

	payload := strings.NewReader(`{"query":"query MyQuery {\n  bitcoin(network: bitcoin) {\n    blocks(date: {after: \"2022-09-19\"}) {\n      timestamp {\n        time\n      }\n      height\n      blockHash\n    }\n  }\n}\n","variables":"{}"}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-API-KEY", "BQYYhiGbQbZhDiVnYwygXjnXdRcwzHyj")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

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
	test := blocks[0].(map[string]interface{})
	fmt.Println(reflect.TypeOf(blocks[0]))
	blockhash := test["blockHash"]
	fmt.Println(blockhash)
}
