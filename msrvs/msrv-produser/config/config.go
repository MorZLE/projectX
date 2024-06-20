package config

import (
	"log"
	"os"
)

func InitConfig() (*Config, error) {
	return parseConfig()
}

type Config struct {
	Rest   Rest   `yaml:"Rest"`
	Broker Broker `yaml:"Broker"`
}

type Broker struct {
	Host string `yaml:"host" env:"RABBITMQ_HOST"`
}
type Rest struct {
	Host string `yaml:"host" env:"REST_HOST"`
}

func parseConfig() (*Config, error) {
	var cnf Config

	cnf.parseEnv()
	cnf.checkEmpty()

	return &cnf, nil
}

func (c *Config) parseEnv() {
	if os.Getenv("RABBITMQ_HOST") != "" {
		c.Broker.Host = os.Getenv("RABBITMQ_HOST")
	}

	if os.Getenv("REST_HOST") != "" {
		c.Rest.Host = os.Getenv("REST_HOST")
	}
}

func (c *Config) checkEmpty() {
	if c.Broker.Host == "" {
		log.Fatalf("BrokerHost is empty")
	}

	if c.Rest.Host == "" {
		c.Rest.Host = ":8080"
	}
}
