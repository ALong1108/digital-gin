package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"time"
)

// Config ..
type Config struct {
	StdConfig `mapstructure:",squash"`

	// connections
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
}

// StdConfig ..
type StdConfig struct {
	Driver   string // "mysql"
	Username string
	Password string
	Protocol string
	Address  string
	Port     uint16
	Dbname   string
	Params   string
}

var dsnFormat = "%s:%s@%s(%s)/%s?%s"

func (m *Config) dsn() string {
	if m.Port > 0 {
		m.Address = fmt.Sprintf("%s:%d", m.Address, m.Port)
	}
	return fmt.Sprintf(dsnFormat, m.Username, m.Password, m.Protocol, m.Address, m.Dbname, m.Params)
}

func (m *Config) mustConnect() *sqlx.DB {
	var db = sqlx.MustConnect(strings.ToLower(m.Driver), m.dsn())
	db.SetMaxOpenConns(m.MaxOpenConns)
	db.SetMaxIdleConns(m.MaxIdleConns)
	db.SetConnMaxLifetime(m.ConnMaxLifetime)
	db.SetConnMaxIdleTime(m.ConnMaxIdleTime)
	return db
}
