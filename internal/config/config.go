package config

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	defaultHttpPort = "80"
)

type (
	Config struct {
		HTTP  HTTPConfig
		MySQL MySQLConfig
	}

	HTTPConfig struct {
		Port string
	}

	MySQLConfig struct {
		URI          string
		DatabaseName string
		User         string
		Password     string
	}
)

func Init(path string) (*Config, error) {
	populateDefaults()

	if err := parseConfigFile(path); err != nil {
		return nil, err
	}

	if err := parseEnv(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	// FIX: Не подгружаются переменные окружения
	// setFromEnv(&cfg)

	return &cfg, nil
}

func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("mysql", &cfg.MySQL); err != nil {
		return err
	}

	return nil
}

func setFromEnv(cfg *Config) {
	cfg.MySQL.URI = viper.GetString("uri")
	cfg.MySQL.User = viper.GetString("user")
	cfg.MySQL.Password = viper.GetString("password")
	log.Info(viper.GetString("password"))
}

func parseEnv() error {
	if err := parseMySQLEnvVariables(); err != nil {
		return err
	}

	return nil
}

func parseMySQLEnvVariables() error {
	viper.SetEnvPrefix("mysql")

	if err := viper.BindEnv("uri"); err != nil {
		return err
	}

	if err := viper.BindEnv("user"); err != nil {
		return err
	}

	return viper.BindEnv("password")
}

func parseConfigFile(filepath string) error {
	path := strings.Split(filepath, "/")
	viper.AddConfigPath(path[0])
	viper.SetConfigName(path[1])

	return viper.ReadInConfig()
}

func populateDefaults() {
	viper.SetDefault("http.port", defaultHttpPort)
}
