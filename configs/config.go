package configs

import (
	"digital-gin/middleware/db"
	"digital-gin/middleware/logger"
	"github.com/spf13/viper"
	"path/filepath"
	"sync"
)

type MyConfig struct {
	AppName string

	HTTP HTTPServer

	Grpc GrpcServer

	GrpcClient GrpcClients

	Logger logger.Config

	DB db.DB

	Redis Redis

	Kafka Kafka

	NebulaGraph NebulaGraph

	Prometheus Prometheus
}

var onceLoadConfig = new(sync.Once)

var myConfig MyConfig

func LoadConfig() {
	onceLoadConfig.Do(func() { loadConfigByEnv(env) })
}

func loadConfigByEnv(env environment) {
	if len(env) == 0 {
		panic("declare env as dev, test or prod")
	}

	loadCfgByPath(filepath.Join(".", "config", string(env)))
}

func loadCfgByPath(path string) {
	viper.SetConfigName("config")
	viper.AddConfigPath(path)
	var err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&myConfig)
	if err != nil {
		panic(err)
	}
}

func Init() error {
	LoadConfig()

	myConfig.Logger.Init()

	myConfig.DB.Init()

	return nil
}
