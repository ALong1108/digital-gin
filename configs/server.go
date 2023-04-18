package configs

import (
	"time"
)

type (
	HTTPServer struct {
		Address      string
		ReadTimeout  time.Duration // time.Millisecond
		WriteTimeout time.Duration // time.Millisecond

		Register `mapstructure:",squash"`
	}

	GrpcServer struct {
		Address        string
		MaxRecvMsgSize int
		MaxSendMsgSize int

		Register `mapstructure:",squash"`
	}

	// Register 服务注册
	Register struct {
		Type           string // static,consul
		Addresses      []string
		ServerName     string
		UpdateInterval int8
		Tags           []string
	}
)

func GetHTTPServer() HTTPServer {
	return myConfig.HTTP
}
