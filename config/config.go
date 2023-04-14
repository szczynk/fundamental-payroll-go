package config

import (
	"fundamental-payroll-go/helper/apperrors"

	"github.com/spf13/viper"
)

type Config struct {
	Debug    bool     `mapstructure:"debug"`
	Port     string   `mapstructure:"port"`
	Storage  string   `mapstructure:"storage"`
	Handler  string   `mapstructure:"handler"`
	Database Database `mapstructure:"db"`
}

type Database struct {
	Driver string `mapstructure:"driver"`
	URL    string `mapstructure:"url"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		_, ok := err.(viper.ConfigFileNotFoundError)
		if ok {
			return nil, apperrors.New(apperrors.ErrEnvNotFound)
		}
		return nil, err
	}

	config := new(Config)
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
