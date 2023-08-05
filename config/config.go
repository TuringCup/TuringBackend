package config

import (
	"os"

	"github.com/spf13/viper"
)

var config *Config

type Config struct {
	System *System `yaml:"system"`
}

type System struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func InitConfig() {
	workdir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workdir)
	viper.AddConfigPath(workdir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
}
