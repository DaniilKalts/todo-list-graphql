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

type DatabaseConfig struct {
	Host     string `mapstructure:"db_host"`
	Port     int    `mapstructure:"db_port"`
	User     string `mapstructure:"db_user"`
	Password string `mapstructure:"db_password"`
	Name     string `mapstructure:"db_name"`
	SSLMode  string `mapstructure:"db_sslmode"`
}

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
}

var cfg Config

func LoadConfig() (*Config, error) {
	v := viper.New()

	v.SetDefault("server.http_port", ":8000")

	v.SetDefault("db_host", "localhost")
	v.SetDefault("db_port", 5432)
	v.SetDefault("db_user", "postgres")
	v.SetDefault("db_password", "1234")
	v.SetDefault("db_name", "spichka")
	v.SetDefault("db_sslmode", "disable")

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
