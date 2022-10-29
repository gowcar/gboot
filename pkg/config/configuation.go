package config

import (
	"github.com/spf13/viper"
	"log"
)

type AppConfig struct {
	Application Application `yaml:"application"`
	Log         Log         `yaml:"log"`
	DB          DB          `yaml:"DB"`
}

type Application struct {
	Name string `yaml:"name"`
	Port int    `yaml:"port"`
}

type Log struct {
	Level  string `yaml:"log_level" mapstructure:"log_level"`
	Folder string `yaml:"folder"`
	File   string `yaml:"file"`
}

type DB struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	Ip       string `yaml:"ip"`
	Port     int    `yaml:"port"`
	MaxConn  int    `yaml:"maxConn" mapstructure:"max_conn"`
	MaxIdle  int    `yaml:"maxIdle" mapstructure:"max_idle"`
	IdleTime int    `yaml:"idleTime" mapstructure:"idle_time"`
	PrintSQL bool   `yaml:"printSQL" mapstructure:"print_sql"`
}

var configHolder *AppConfig

func InitConfig(config_file string) {
	viper.SetConfigName(config_file)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read config failed: ", err)
	}
	if err != nil {
		panic(err)
	}
	configHolder = &AppConfig{}
	err = viper.Unmarshal(&configHolder)
	if err != nil {
		panic(err)
	}
}

func ConfigGet(key string) any {
	return viper.Get(key)
}

func Config() *AppConfig {
	return configHolder
}
