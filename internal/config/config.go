package config

import (
	"fmt"
	"os"
	"encoding/json"
)

const configFileName = ".gatorconfig.json";

type Config struct {
	DbUrl string `json:"db_url"`
	CurrentUserName string `json:"current_user_name,omitempty"`
	ConnectionString string `json:"connection_string"`
}

func GetConfigFilePath() (string, error) {
	name, err := os.UserHomeDir(); if err != nil {
		return "", fmt.Errorf("can't find home directory + configFileName concatenation");
	}
	name += "/";
	name += configFileName;
	// _, err := os.Stat("~/" + configFileName); if err != nil {
	// 	return "", errors.New("Config file path not found in home directory");
	// }
	return name, nil;
}

func Read() (*Config, error) {
	fileName, err := GetConfigFilePath(); if err != nil {
		return nil, fmt.Errorf("config file path error: %s", err)
	}
	data, err := os.ReadFile(fileName); if err != nil {
		return nil, fmt.Errorf("config file read error: %s", err);
	}
	var parsedConfig Config
	if err := json.Unmarshal(data, &parsedConfig); err != nil {
		return nil, fmt.Errorf("config file unmarshaling json: %s", err)
	}
	return &parsedConfig, nil
}

func (c *Config) SetUser(UserName string) error {
	if len(UserName) == 0 {
		return fmt.Errorf("user name cannot be empty")
	}
	c.CurrentUserName = UserName;
	err := write(c); if err != nil {
		return err
	}
	return nil
}

func write(cfg *Config) error {
	fileName, err := GetConfigFilePath(); if err != nil {
		return fmt.Errorf("config file path error: %s", err)
	}
	data, err := json.Marshal(cfg); if err != nil {
		return fmt.Errorf("couldn't marshal data to json: %s", err);
	}
	permissions := os.FileMode(0644);
	err3 := os.WriteFile(fileName, data, permissions) 
	if err3 != nil {
		return fmt.Errorf("error writing config file: %s", err)
	}	
	return nil
}

