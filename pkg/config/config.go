package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig
	Jwt      JwtConfig
	Database DatabaseConfig
}

type AppConfig struct {
	Port string `mapstructure:"port"`
	Env  string `mapstructure:"env"`
}

type JwtConfig struct {
	Secret string `mapstructure:"secret"`
	Expire int    `mapstructure:"expire"`
}

type DatabaseConfig struct {
	Path string `mapstructure:"path"`
}

var Current Config
var IsProduction bool

func Init() error {
	viper.SetConfigFile("config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&Current); err != nil {
		return err
	}

	IsProduction = Current.App.Env == "production"

	return nil
}
