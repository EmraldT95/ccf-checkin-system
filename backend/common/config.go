package common

import (
	"fmt"

	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type Config struct {
	Database DatabaseConfig `yaml:"database"`
}

// Parses the config file and returns the config
func ParseConfigFile() (*Config, error) {
	// Read the config file
	err := viper.ReadInConfig()
	if err != nil {
		// Handle the case where config file is not found
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, fmt.Errorf("No config file found, using defaults")
		}

		return nil, fmt.Errorf("Error reading config file: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	if config.Database == (DatabaseConfig{}) {
		return nil, fmt.Errorf("missing database config")
	}

	return &config, nil
}
