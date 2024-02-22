package options

import (
	"errors"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Options struct {
	//API settings
	APIAddr string
}

var (
	errReadConfig      = errors.New("read config error")
	errEmptyConfigName = errors.New("configName is empty")
)

var options *Options

func LoadConfig(configName string) (*Options, error) {
	log.Info("Try to load configuration file...")

	if configName == "" {
		return nil, errEmptyConfigName
	}

	viper.SetConfigName(configName)
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Warnf("Configuration file not found")
		return nil, errReadConfig
	}
	options = &Options{
		APIAddr: viper.GetString("api.addr"),
	}

	return options, nil
}
