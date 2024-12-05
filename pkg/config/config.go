package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	ChainRPCs map[uint64]string `json:"chain_rpcs"`
	GRPCPort  string            `json:"grpc_port"`
}

func LoadConfig(configPath string) (*Config, error) {
	config := &Config{}

	file, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	if err := json.Unmarshal(file, config); err != nil {
		return nil, fmt.Errorf("error parsing config file: %w", err)
	}

	return config, nil
}
