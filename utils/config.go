package utils

import (
	"encoding/json"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string `json:"port"`
	Host string `json:"host"`
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	config := new(Config)
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}
	return config, nil
}

func ParseEnv(path string) error {
	return godotenv.Load(path)
}

func LoadAllowedTypes(dest string) ([]string, error) {
	filename := "config/mime.json"
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := make(map[string][]string)
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}
	return config[dest], nil
}
