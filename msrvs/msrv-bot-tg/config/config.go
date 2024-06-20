package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"strconv"
	"strings"
)

func InitConfig() (*Config, error) {
	return parseConfig()
}

type Config struct {
	dev    string `yaml:"Dev"`
	Broker Broker `yaml:"Broker"`
	Bot    Bot    `yaml:"Bot"`
}
type Broker struct {
	Host string `yaml:"host" env:"RABBITMQ_HOST"`
}

type Bot struct {
	Token      string   `yaml:"token" env:"BOT_TOKEN"`
	TimeUpdate int      `yaml:"timeUpdate" env:"BOT_TIME_UPDATE"`
	Admins     []string `yaml:"admins" env:"BOT_ADMINS"`
}

func parseConfig() (*Config, error) {
	var cnf Config
	cnfFile, err := os.ReadFile("msrvs/msrv-bot-tg/config/config.yaml")
	if err == nil {
		err = yaml.Unmarshal(cnfFile, &cnf)
		if err != nil || cnf.dev != "docker" {
			cnf.parseEnv()
		}
	}

	cnf.checkEmpty()

	return &cnf, nil
}

func (c *Config) parseEnv() {
	c.Broker.Host = os.Getenv("RABBITMQ_HOST")
	c.Bot.Token = os.Getenv("BOT_TOKEN")

	if os.Getenv("BOT_TIME_UPDATE") != "" {
		n, err := strconv.Atoi(os.Getenv("BOT_TIME_UPDATE"))
		if err != nil && c.Bot.TimeUpdate == 0 {
			c.Bot.TimeUpdate = 2
		} else {
			c.Bot.TimeUpdate = n
		}
	}

	c.Bot.Admins = strings.Split(os.Getenv("ADMIN"), ",")
}

func (c *Config) checkEmpty() {
	if c.Bot.Token == "" {
		log.Fatalf("BotToken is empty")
	}

	if c.Broker.Host == "" {
		log.Fatalf("BrokerHost is empty")
	}

	if len(c.Bot.Admins) == 0 {
		log.Fatalf("Admins is empty")
	}
}
