package config

import (
	"encoding/json"
	"fmt"
	"os"
	"wb/rest-api/pkg/logging"
)

type Config struct {
	DB     Database `json:"DB"`
	Listen Server   `json:"listen"`
}

type Database struct {
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"DBName"`
	SSLMode  string `json:"SSLMode"`
}

type Server struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func GetConfig(cfgPath string, logger *logging.Logger) (*Config, error) {
	logger.Infof("get config from: %s", cfgPath)
	cfg := &Config{}

	bytesCfg, err := os.ReadFile(cfgPath)
	if err != nil {
		logger.Warningf("unable to read cfg: %v", err)
		return nil, fmt.Errorf("unable to read cfg file: %v", err)
	}

	err = json.Unmarshal(bytesCfg, &cfg)
	if err != nil {
		logger.Warningf("unable to convert cfg to model: %v", err)
		return nil, fmt.Errorf("unable to convert cfg to model: %v", err)
	}

	return cfg, nil
}
