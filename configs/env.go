package configs

import (
	"github.com/spf13/viper"
)

type environment string

const (
	dev  environment = "dev"
	test environment = "test"
	prod environment = "prod"
)

var env environment

func init() {
	err := viper.BindEnv("env")
	if err != nil {
		panic(err)
	}
	env = environment(viper.GetString("env"))
}

func IsDevEnv() bool {
	return env == dev
}

func IsTestEnv() bool {
	return env == test
}

func IsProdEnv() bool {
	return env == prod
}
