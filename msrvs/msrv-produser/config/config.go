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
	RestHost   string `yaml:"RestHost"`
	BrokerHost string `yaml:"BrokerHost"`
}

func parseConfig() (*Config, error) {
	cnfFile, err := os.ReadFile("msrvs/msrv-produser/config/config.yaml") //TODO сделать кросс
	if err != nil {
		return nil, err
	}

	var cnf Config
	err = yaml.Unmarshal(cnfFile, &cnf)
	if err != nil {
		return nil, err
	}

	if os.Getenv("PORT") != "" {
		cnf.RestHost = ":" + os.Getenv("PORT")
	}
	if os.Getenv("RABBITMQ_HOST") != "" {
		cnf.BrokerHost = os.Getenv("RABBITMQ_HOST")
	}

	if cnf.RestHost == "" {
		cnf.RestHost = ":8080"
	}
	if cnf.BrokerHost == "" {
		log.Fatalf("BrokerHost is empty")
		//cnf.BrokerHost = "amqp://rmuser:rmpassword@localhost:5672/"
	}

	return &cnf, nil
}
