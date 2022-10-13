package config

import (
	"ginblog/storage"
	"github.com/pkg/errors"
)

// Config defines the configuration structure
type Config struct {
	General struct {
		LogLevel   int    `mapstructure:"log_level"`
		Port       int    `mapstructure:"port"`
		Mode       string `mapstructure:"mode"`
		UploadPath string `mapstructure:"uploadpath"`
	} `mapstructure:"general"`

	MySQL struct {
		DSN         string `mapstructure:"dsn"`
		AutoMigrate bool   `mapstructure:"auto_migrate"`
		DB          *storage.DBLogger
	} `mapstructure:"mysql"`

	BlogConf struct {
		Data map[string]interface{} `mapstructure:"data"`
	} `mapstructure:"blogconf"`
}

var C Config

func (c *Config) String(key string) (string, error) {
	if v, ok := c.BlogConf.Data[key].(string); ok {
		return v, nil
	}
	return "", errors.New(key + " value not fond")

}

func (c *Config) Bool(key string) (bool, error) {
	if v, ok := c.BlogConf.Data[key].(bool); ok {
		return v, nil
	}
	return false, errors.New(key + " value not fond")
}

func (c *Config) Int(key string) (int, error) {
	if v, ok := c.BlogConf.Data[key]; ok {
		switch v.(type) {
		case int:
			return v.(int), nil
		case int8:
			return int(v.(int8)), nil
		case int16:
			return int(v.(int16)), nil
		case int32:
			return int(v.(int32)), nil
		case int64:
			return int(v.(int64)), nil
		default:
			return 0, errors.New("not interger")
		}
	}
	return 0, errors.New(key + " value not fond")
}

func (c *Config) Int64(key string) (int64, error) {
	if v, ok := c.BlogConf.Data[key]; ok {
		switch v.(type) {
		case int:
			return int64(v.(int)), nil
		case int8:
			return int64(v.(int8)), nil
		case int16:
			return int64(v.(int16)), nil
		case int32:
			return int64(v.(int32)), nil
		case int64:
			return v.(int64), nil
		default:
			return 0, errors.New("value is not interger type")
		}
	}
	return 0, errors.New(key + " value not fond")
}

func (c *Config) Float(key string) (float64, error) {
	if v, ok := c.BlogConf.Data[key]; ok {
		switch v.(type) {
		case float32:
			return float64(v.(float32)), nil
		case float64:
			return v.(float64), nil
		default:
			return 0.0, errors.New("value is not float type")
		}
	}
	return 0.0, errors.New(key + " value not fond")
}

func (c *Config) DIY(key string) (interface{}, error) {
	if v, ok := c.BlogConf.Data[key]; ok {
		return v, nil
	}
	return nil, errors.New(key + " value not fond")
}
