package configs

import "time"

// NebulaGraph ..
type NebulaGraph struct {
	Address  string
	Port     int
	Username string
	Password string

	TimeOut         time.Duration // Socket timeout and Socket connection timeout, unit: seconds
	IdleTime        time.Duration // The idleTime of the connection, unit: seconds
	MaxConnPoolSize int           // The max connections in pool for all addresses
	MinConnPoolSize int           // The min connections in pool for all addresses

	KeepAlive time.Duration
}
