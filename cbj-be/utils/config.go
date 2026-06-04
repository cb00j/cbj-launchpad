package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

type Config struct {
	MySQL  MySQLConfig
	Signer SignerConfig
}

type MySQLConfig struct {
	IP       string `ini:"ip"`
	Port     string `ini:"port"`
	User     string `ini:"user"`
	Password string `ini:"password"`
	Database string `ini:"database"`
}

type SignerConfig struct {
	PrivateKey string `ini:"private_key"`
}

func LoadConfig(path string) (*Config, error) {
	config, err := ini.Load(path)
	if err != nil {
		return nil, fmt.Errorf("load config %s: %w", path, err)
	}

	cfg := &Config{}

	if err := config.Section("mysql").MapTo(&cfg.MySQL); err != nil {
		return nil, fmt.Errorf("map mysql config: %w", err)
	}
	if err := config.Section("signer").MapTo(&cfg.Signer); err != nil {
		return nil, fmt.Errorf("map signer config: %w", err)
	}

	if cfg.Signer.PrivateKey == "" {
		return nil, fmt.Errorf("private key is required")
	}

	return cfg, nil
}
