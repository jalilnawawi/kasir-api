package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}

func LoadConfig() (*Config, error) {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		err := viper.ReadInConfig()
		if err != nil {
			return nil, fmt.Errorf("Error loading config file: %w", err)
		}
	}

	config := &Config{
		Port:   viper.GetString("APP_PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}

	return config, nil
}
