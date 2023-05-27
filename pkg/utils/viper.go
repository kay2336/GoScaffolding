package utils

import (
	"awesomeProject/model/config"
	"github.com/spf13/viper"
)

func GetMysqlByViper() config.ConnMysql {
	v := viper.New()
	v.AddConfigPath("config")
	v.SetConfigType("yaml")

	var connMysql config.ConnMysql
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = v.UnmarshalKey("mysql", &connMysql)
	if err != nil {
		panic(err)
	}
	return connMysql
}

func GetRouterByViper() config.ConnGin {
	v := viper.New()
	v.AddConfigPath("config")
	v.SetConfigType("yaml")

	var connGin config.ConnGin
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = v.UnmarshalKey("router", &connGin)
	if err != nil {
		panic(err)
	}
	return connGin
}
