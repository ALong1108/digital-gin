package configs

import "time"

// Server ..
type Server struct {
	Address string

	// 证书
	CertFile string
	KeyFile  string
	Phrase   string

	//服务注册配置
	//服务注册类型，目前可选static、consul，默认是static
	ResolverType string
	//服务注册的地址，如consul、etcd地址
	ResolverAddresses []string
	//服务名，必须符合证书的域名规则
	ServerName string
	//隔多久检查一次
	UpdateInterval int64
	//指定注册的网卡地址，或者在上方的Address里指定ip
	Interface string
	Tags      []string
}

// HTTPServer ..
type HTTPServer struct {
	Server `mapstructure:",squash"`
	//并发数量限制
	Concurrence uint

	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// GrpcServer ..
type GrpcServer struct {
	Server `mapstructure:",squash"`

	MaxRecvMsgSize int
	MaxSendMsgSize int
}
