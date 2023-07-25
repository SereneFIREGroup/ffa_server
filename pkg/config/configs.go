package config

import (
	"fmt"

	"github.com/juju/errors"
	"github.com/spf13/viper"
)

const (
	ProductNameZH = "悠然度日"
	ProductNameEN = "SereneWealth"
)

var config map[string]interface{}

func LoadConfig() error {
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("config/")
	err := viper.ReadInConfig()
	if err != nil {
		return errors.Trace(err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return errors.Trace(err)
	} //将配置文件绑定到config上
	return nil
}

func StringOrPanic(key string) string {
	v, err := StringOrError(key)
	if err != nil {
		panic(err)
	}
	return v
}

func StringOrError(key string) (value string, err error) {
	value, found := getString(key, "")
	if !found {
		err = fmt.Errorf("%s is not configured", key)
	}
	return
}

// String 根据某个 key 获取配置值，值类型为：string
func String(key string, defaultValue string) (value string) {
	value, _ = getString(key, defaultValue)
	return
}

func getString(key string, defaultValue string) (value string, found bool) {
	v, found := config[key]
	if !found {
		return defaultValue, false
	}
	if value, ok := v.(string); ok {
		return value, ok
	}
	return defaultValue, false
}

// Bool 根据某个 key 获取配置值，值类型为：bool
func Bool(key string, defaultValue bool) (value bool) {
	value, _ = getBool(key, defaultValue)
	return
}

func getBool(key string, defaultValue bool) (value bool, found bool) {
	v, found := config[key]
	if !found {
		value = defaultValue
		return
	}

	if b, ok := v.(bool); ok {
		return b, ok
	} else if i, ok := v.(int); ok {
		if i == 1 {
			return true, true
		} else if i == 0 {
			return false, true
		} else {
			return defaultValue, false
		}
	} else {
		return defaultValue, false
	}
}

// Int 根据某个 key 获取配置值，值类型为：int
func Int(key string, defaultValue int) (value int) {
	value, _ = getInt(key, defaultValue)
	return
}

func getInt(key string, defaultValue int) (value int, found bool) {
	v, found := config[key]
	if !found {
		value = defaultValue
		return
	}

	if v64, found := v.(float64); found {
		return int(v64), found
	} else {
		return defaultValue, false
	}
}

// Int64 根据某个 key 获取配置值，值类型为：int64
func Int64(key string, defaultValue int64) (value int64, found bool) {
	v, found := config[key]
	if !found {
		value = defaultValue
		return
	}

	if v64, found := v.(float64); found {
		return int64(v64), found
	} else {
		return defaultValue, false
	}
}

func IntOrPanic(key string) int {
	v, err := IntOrError(key)
	if err != nil {
		panic(err)
	}
	return v
}

func IntOrError(key string) (value int, err error) {
	value, found := getInt(key, 0)
	if !found {
		err = fmt.Errorf("%s is not configured", key)
	}
	return
}
