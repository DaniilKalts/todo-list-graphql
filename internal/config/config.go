package config

import (
	"errors"
	"fmt"

	"github.com/goforj/godump"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Port string `mapstructure:"http_port"`
}

type Config struct {
	Server ServerConfig `mapstructure:"server"`
}

var cfg Config

func LoadConfig() (*Config, error) {
	v := viper.New()

	v.SetDefault("server.http_port", ":8000")

	v.SetConfigName("config.local")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")

	if err := v.ReadInConfig(); err != nil {
		if !errors.Is(err, viper.ConfigFileNotFoundError{}) {
			return nil, fmt.Errorf("failed to read config: %w", err)
		}
		fmt.Println("Loaded configuration from default values")
	}

	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}
	if v.ConfigFileUsed() != "" {
		fmt.Printf("Loaded configuration from: %s\n", v.ConfigFileUsed())
	}

	godump.Dump(cfg)

	return &cfg, nil
}
