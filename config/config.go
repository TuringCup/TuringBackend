package config

import (
	"os"

	"github.com/spf13/viper"
)

var Conf *Config

type Config struct {
	System     *System     `yaml:"system"`
	DB         *DB         `yaml:"db"`
	Skywalking *Skywalking `yaml:"skywalking"`
	SES        *SES        `yaml:"ses"`
}

type System struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type DB struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DbName   string `yaml:"dbName"`
	UserName string `yaml:"userName"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
}

type Skywalking struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type SES struct {
	SecretID  string `yaml:"id"`
	SecretKey string `yaml:"key"`
}

func InitConfig(path string) {
	workdir := path
	if path == "" {
		workdir, _ = os.Getwd()
	}
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workdir)
	viper.AddConfigPath(workdir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&Conf)
	if err != nil {
		panic(err)
	}
}
