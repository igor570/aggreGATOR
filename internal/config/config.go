package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const gatorConfigFile = ".gatorconfig.json"

type Config struct {
	DatabaseURL string `json:"db_url"`
	UserName    string `json:"current_user_name"`
}

func NewConfig() *Config {
	return &Config{
		DatabaseURL: "",
	}
}

func (c *Config) ReadConfig() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	configPath := filepath.Join(home, gatorConfigFile)

	jsonData, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(jsonData, c)
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) SetUser(userName string) error {
	if userName == "" {
		return fmt.Errorf("Username cannot be empty")
	}

	c.UserName = userName

	data, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		return err
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	configPath := filepath.Join(home, gatorConfigFile)

	return os.WriteFile(configPath, data, 0666)
}
