package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	Origin string `mapstructure:"CLIENT_ORIGIN"`
	DBUri  string `mapstructure:"MONGODB_URI"`
	DBName string `mapstructure:"MONGODB_DBNAME"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
