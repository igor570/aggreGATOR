package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

type AuthUser struct {
	LoggedUser string `json:"loggedUser"`
}

func CheckAuth() (*AuthUser, error) {
	var a AuthUser
	configPath := "../auth.json" // points to auth.json in the project root

	jsonData, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonData, &a)
	if err != nil {
		return nil, err
	}

	if a.LoggedUser == "" {
		return nil, fmt.Errorf("No user is logged in. Please log in to write a user to the file.")
	}

	return &a, nil
}
