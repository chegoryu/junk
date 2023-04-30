package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Port int `json:"port"`
}

func LoadConfig() Config {
	var configPath string
	{
		flag.StringVar(&configPath, "config", "config.json", "Config file")

		configFlag := flag.Lookup("config")
		flag.Var(configFlag.Value, "c", fmt.Sprintf("Alias to %s", configFlag.Name))
	}

	flag.Parse()

	configFile, err := os.Open(configPath)
	if err != nil {
		log.Fatalf("Failed to open config file: %v", err)
	}
	defer configFile.Close()

	configString, err := ioutil.ReadAll(configFile)
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	var config Config
	err = json.Unmarshal(configString, &config)
	if err != nil {
		log.Fatalf("Failed to parse config: %v", err)
	}

	return config
}
