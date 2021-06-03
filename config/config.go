package config

import "github.com/spf13/viper"

func Load() (Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}
	config := Config{}
	if err := viper.Unmarshal(&config); err != nil {
		return Config{}, err
	}
	return config, nil
}

type Config struct {
	Database Database
	Server   Server
}

type Database struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
	SSLMode  string
}

type Server struct {
	Port string
}
