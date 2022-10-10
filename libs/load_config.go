package libs

import (
	"blockdata-crawling/src/types"
	"encoding/json"
	"fmt"
	"os"
)

func LoadConfig() types.Config {
	var config types.Config
	file, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()
	json.NewDecoder(file).Decode(&config)
	return config
}
