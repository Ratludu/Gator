package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configPath = "/.gatorconfig.json"

func getHomeDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return homeDir, nil

}
func Read() Config {

	homeDir, err := getHomeDir()
	if err != nil {
		fmt.Println("Error: could not get home directory", err)
		return Config{}
	}

	filePath := filepath.Join(homeDir, configPath)

	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error: could not read json data", err)
		return Config{}
	}

	var config Config
	err = json.Unmarshal(jsonData, &config)
	if err != nil {
		fmt.Println("Error: could not unpack json", err)
	}

	return config
}

func (c *Config) SetUser(username string) {

	c.CurrentUserName = username

	homeDir, err := getHomeDir()
	if err != nil {
		log.Fatalf("Could not get home directory: %v", err)
	}

	filePath := filepath.Join(homeDir, configPath)

	jsonData, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		log.Fatalf("Error unable to marshal config struct: %v", err)
	}

	err = os.WriteFile(filePath, jsonData, 0666)
	if err != nil {
		log.Fatalf("Cannot write to file: %v", err)
	}
}

