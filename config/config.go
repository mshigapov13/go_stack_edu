package config

import (
	"encoding/json"
	"flag"
	"io"
	"log"
	"os"
)

type Config struct {
	Bot bot `json:"bot"`
}

type bot struct {
	Platform string `json:"platform"`
	Name     string `json:"name"`
	Token    string `json:"token"`
}

func readConfigJSON(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg Config
	byteValue, _ := io.ReadAll(file)
	err = json.Unmarshal(byteValue, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func LoadConfigJson() *Config {
	configJSON := flag.String("bot", "config.json", "Defines configuration file option")
	flag.Parse()

	cfg, err := readConfigJSON(*configJSON)
	if err != nil {
		log.Fatal(err)
	}
	return cfg
}
