package config

import (
	"gopkg.in/yaml.v3"
	"log"
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

	if os.Getenv("RABBITMQ_HOST") != "" {
		cnf.BrokerHost = os.Getenv("RABBITMQ_HOST")
	}

	if cnf.BrokerHost == "" {
		log.Fatalf("BrokerHost is empty")
		//cnf.BrokerHost = "amqp://rmuser:rmpassword@localhost:5672/"
	}
	return &cnf, nil
}
