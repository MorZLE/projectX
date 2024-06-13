package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

func InitConfig() (*Config, error) {
	return parseConfig()
}

type Config struct {
	BrokerHost string `yaml:"BrokerHost"`
}

func parseConfig() (*Config, error) {
	cnfFile, err := os.ReadFile("msrvs/msrv-bot-tg/config/config.yaml") //TODO сделать кросс
	if err != nil {
		return nil, err
	}

	var cnf Config
	err = yaml.Unmarshal(cnfFile, &cnf)
	if err != nil {
		return nil, err
	}

	return &cnf, nil
}
