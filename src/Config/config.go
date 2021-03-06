package Config

import (
	"encoding/json"
	"os"
	"fmt"
)

type Config struct {
	ClientID	string	`json:clientId`
	Key			string	`json:"key"`
}

func LoadConfiguration(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}