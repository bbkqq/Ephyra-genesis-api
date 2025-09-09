package mysql

import (
	"Ephyra-genesis-api/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

const (
	defaultMaxOpenConns    = 128
	defaultConnMaxIdleTime = 300 * time.Second
	defaultConnMaxLifetime = 300 * time.Second
	defaultTimeout         = 5 * time.Second
	defaultReadTimeout     = 5 * time.Second
	defaultWriteTimeout    = 5 * time.Second
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	DB, err = gorm.Open(mysql.Open(conf.GetConf().MySQL.DSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	mysqlDB, err := DB.DB()
	if err != nil {
		panic(err)
	}

	err = mysqlDB.Ping()
	if err != nil {
		panic(err)
	}

	maxOpenConns := defaultMaxOpenConns

	mysqlDB.SetMaxOpenConns(maxOpenConns)

	mysqlDB.SetMaxIdleConns(maxOpenConns / 2)

	connMaxIdleTime := defaultConnMaxIdleTime
	mysqlDB.SetConnMaxIdleTime(connMaxIdleTime)

	connMaxLifetime := defaultConnMaxLifetime
	mysqlDB.SetConnMaxLifetime(connMaxLifetime)
}
