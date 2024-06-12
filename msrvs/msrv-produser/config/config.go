package config

func InitConfig() (*Config, error) {
	return parseConfig()
}

func parseConfig() (*Config, error) {
	return &Config{}, nil
}

type Config struct {
}
