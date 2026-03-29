package config

import (
	"strings"

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
	Secret     string `mapstructure:"secret"`
	Expiration int    `mapstructure:"expiration"`
}

type DatabaseConfig struct {
	Driver string `mapstructure:"driver"`
	Path   string `mapstructure:"path"`
	DSN    string `mapstructure:"dsn"`
}

func (c DatabaseConfig) DriverName() string {
	driver := strings.TrimSpace(strings.ToLower(c.Driver))
	switch driver {
	case "", "sqlite", "sqlite3":
		return "sqlite"
	case "postgres", "postgresql":
		return "postgres"
	case "mysql":
		return "mysql"
	default:
		return driver
	}
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
