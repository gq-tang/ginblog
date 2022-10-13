package test

import (
	"ginblog/config"
	"github.com/spf13/viper"
)

func LoadConfig() (config.Config, error) {
	var c config.Config
	viper.SetConfigName("gin-blog")
	viper.AddConfigPath("../blog/")
	if err := viper.ReadInConfig(); err != nil {
		return c, err
	}
	if err := viper.Unmarshal(&c); err != nil {
		return c, err
	}
	return c, nil
}
