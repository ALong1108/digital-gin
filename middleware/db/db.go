package db

import (
	"github.com/jmoiron/sqlx"
	"sync"
)

// DB ..
type DB struct {
	Digital    Config
	Etl        Config
	Statistics Config
	Artery     Config
}

var (
	initOnce   = new(sync.Once)
	closeOnce  = new(sync.Once)
	digital    *sqlx.DB
	etl        *sqlx.DB
	statistics *sqlx.DB
	artery     *sqlx.DB
)

func (db *DB) Init() {
	initOnce.Do(func() {
		digital = db.Digital.mustConnect()
		etl = db.Etl.mustConnect()
		statistics = db.Statistics.mustConnect()
		artery = db.Artery.mustConnect()
	})
}

func Close() {
	closeOnce.Do(func() {
		if digital != nil {
			digital.Close()
		}
		if etl != nil {
			etl.Close()
		}
		if statistics != nil {
			statistics.Close()
		}
		if artery != nil {
			artery.Close()
		}
	})
}

func Digital() *sqlx.DB {
	return digital
}

func Etl() *sqlx.DB {
	return etl
}

func Statistics() *sqlx.DB {
	return statistics
}

func Artery() *sqlx.DB {
	return artery
}
