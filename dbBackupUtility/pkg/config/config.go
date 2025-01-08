package config

import (
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	DBMS string `yaml:"dbms"`
	DBConfig struct{
		Host string `yaml:"host"`
		Port int `yaml:"port"`
		User string `yaml:"user"`
		Password string `yaml:"password"`
		DBName string `yaml:"dbname"`
	} `yaml:"dbConfig"`
	BackupStorage string `yaml:"backupStorage"`
} 

var (
	config *Config
)

func GetConfig(configFilePath string) *Config {
	if config == nil {
		config = loadConfig(configFilePath)
	}
	return config
}

func loadConfig(configFilePath string) *Config {
	if configFilePath == "" {
		log.Fatalf("Error reading config file path: %v", configFilePath)
		return nil
	}
	file, err := os.Open(configFilePath)
	if err != nil {
		log.Fatalf("Error opening config file: %v", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatalf("Error unmarshalling config data: %v", err)
	}

	return &cfg
}