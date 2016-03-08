package main

import (
	"encoding/json"
	"flag"
	"os"

	"github.com/flebel/embassy/ambassadors"
	"github.com/flebel/embassy/embassyd"
)

type Configuration struct {
	Ambassadors []config.Ambassador
	Listen      string
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
	configFile := flag.String("conf", "config.json", "Configuration file")
	flag.Parse()
	config := loadConfig(*configFile)
	embassyd.StartNewEmbassyD(config.Ambassadors, config.Listen)
}
