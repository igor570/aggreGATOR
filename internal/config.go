package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const gaterConfig = "../gaterconfig.json"

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

	jsonData, err := os.ReadFile(gaterConfig)

	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(jsonData, c)

	if err != nil {
		return err
	}

	fmt.Printf("Created struct with properties: %+v\n", c)

	return nil
}

func (c *Config) SetUser(userName string) error {
	if userName == "" {
		return fmt.Errorf("Username cannot be empty")
	}

	c.UserName = userName
	c.DatabaseURL = "placeholderURL"

	data, err := json.MarshalIndent(c, "", " ")

	if err != nil {
		return err
	}

	os.WriteFile(gaterConfig, data, 0666)

	return nil
}
