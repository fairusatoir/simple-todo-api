package config

import (
	"github.com/spf13/viper"
)

func InitializeAppConfig() error {
	viper.SetConfigFile("config.yml")
	viper.AddConfigPath(".")
	viper.AllowEmptyEnv(true)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func AppName() string {
	v := viper.GetString("app.name")
	if len(v) == 0 {
		return "todo.service"
	}
	return v
}

func AppVersion() string {
	v := viper.GetString("app.version")
	if len(v) == 0 {
		return "1.0.0"
	}
	return v
}

func AppEnv() string {
	return viper.GetString("app.environment")
}

func AppDebug() bool {
	v := viper.GetBool("app.debug")
	// v: = true
	return v
}

func DSDatamasterUrl() string {
	return viper.GetString("datasource.datamaster.url")
}

func DSDatamasterDriver() string {
	return viper.GetString("datasource.datamaster.driver")
}
