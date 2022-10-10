package libs

import (
	"blockdata-crawling/src/types"
	"encoding/json"
	"fmt"
	"os"
)

func LoadDataSourceConfig() types.DataSource {
	var datasource types.DataSource
	file, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()
	json.NewDecoder(file).Decode(&datasource)
	return datasource
}
