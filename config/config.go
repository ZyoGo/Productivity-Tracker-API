package config

import (
	"sync"

	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

type AppConfig struct {
	// App
	App struct {
		Name string `toml:"name"`
		Env  string `toml:"env"`
		Port int    `toml:"port"`
	} `toml:"App"`
	// Database
	Database struct {
		Driver   string `toml:"driver"`
		Host     string `toml:"host"`
		Port     int    `toml:"port"`
		Username string `toml:"username"`
		Password string `toml:"password"`
		Name     string `toml:"name"`
		DBUrl    string `toml:"dburl"`
	} `toml:"Database"`
	// Logger
	Logger struct {
		Driver   string `toml:"driver"`
		Host     string `toml:"host"`
		Port     int    `toml:"port"`
		Password string `toml:"password"`
		Dbnumber int    `toml:"dbnumber"`
	} `toml:"Logger"`
}

// mutex for singleton
var lock = &sync.Mutex{}
var appConfig *AppConfig

// GetConfig return singleton
func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

// initConfig return singleton
func initConfig() *AppConfig {
	viper.SetConfigType("toml")
	viper.SetConfigName("config")
	viper.AddConfigPath("./config/")

	if err := viper.ReadInConfig(); err != nil {
		log.Info("error to read file config ", err)
	}

	var finalConfig AppConfig
	err := viper.Unmarshal(&finalConfig)
	if err != nil {
		log.Info("Failed to extract config")
	}

	return &finalConfig
}
