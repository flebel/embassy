package main

import (
	"encoding/json"
	"os"

	"github.com/flebel/embassy/embassyd"
)

type Configuration struct {
	Listen string
}

func loadConfig(filename string) Configuration {
	file, _ := os.Open(filename)
	decoder := json.NewDecoder(file)
	config := Configuration{}
	err := decoder.Decode(&config)
	if err != nil {
		panic("Error: " + err.Error())
	}
	return config
}

func main() {
	config := loadConfig("config.json")
	embassyd.StartNewEmbassyD(config.Listen)
}
