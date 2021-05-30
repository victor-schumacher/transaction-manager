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

type Config struct{
	Db  Database `mapstructure:"database"`
}

type Database struct {

}
