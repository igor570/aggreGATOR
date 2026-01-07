package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBUrl string `json:"db_url"`
	User  string `json:"current_user_name"`
}

func (c *Config) SetUser(name string) error {
	if len(name) == 0 {
		return errors.New("Name cannot be empty")
	}

	// Update the user field in the config
	c.User = name

	filePath := GetGatorConfigDir()

	// Marshal the entire config struct
	contents, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling json: %w", err)
	}

	// Write the config to file
	err = os.WriteFile(filePath, contents, 0644)
	if err != nil {
		return fmt.Errorf("error writing config file: %w", err)
	}

	return nil
}

func Read() (Config, error) {
	filePath := GetGatorConfigDir()

	fileContents, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, fmt.Errorf("error reading config file: %w", err)
	}

	var cfg Config

	// decode the json into cfg so we can get the postgres connection string
	err = json.Unmarshal(fileContents, &cfg)
	if err != nil {
		return cfg, errors.New("Error unmarshelling json")
	}

	return cfg, nil
}

func GetGatorConfigDir() string {
	homeDir, _ := os.UserHomeDir()
	filePath := filepath.Join(homeDir, configFileName)
	return filePath
}
