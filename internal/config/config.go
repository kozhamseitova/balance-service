package config

import "github.com/spf13/viper"

type (
	Config struct {
		App struct {
			Name        string `mapstructure:"NAME"`
			LogLevel    string `mapstructure:"LOG_LEVEL" default:"DEBUG"`
			Environment string `mapstructure:"ENVIRONMENT" default:"DEV"`
			Port        string `mapstructure:"PORT"`
		}
		Database struct {
			DSN string `mapstructure:"DSN"`
		}
	}
)

func New() (*Config, error) {
	var config Config

	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}