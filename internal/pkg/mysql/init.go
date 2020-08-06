package mysql

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

var (
	mysqlDB *sqlx.DB
	err     error
)

func InitMysql(address string, idleConn, openConn int, maxLifeTime time.Duration) {
	mysqlDB, err = sqlx.Open("mysql", address)
	// 打印日志
	if err != nil {
		panic(err)
	}
	err = mysqlDB.Ping()
	if err != nil {
		panic(err)
	}
	mysqlDB.SetMaxIdleConns(idleConn)
	mysqlDB.SetMaxOpenConns(openConn)
	mysqlDB.SetConnMaxLifetime(maxLifeTime)
	log.Println("Mysql is Collection!!!")
}
