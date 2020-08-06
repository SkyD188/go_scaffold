package mysql

import (
	"go_scaffold/config"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

var (
	mysqlDB *sqlx.DB
	err     error
)

func InitMysql() {
	mysqlDB, err = sqlx.Open("mysql", config.GetConf().Mysql.Address)
	// 打印日志
	if err != nil {
		panic(err)
	}
	err = mysqlDB.Ping()
	if err != nil {
		panic(err)
	}
	mysqlDB.SetMaxIdleConns(config.GetConf().Mysql.MaxIdleConn)
	mysqlDB.SetMaxOpenConns(config.GetConf().Mysql.MaxOpenConn)
	mysqlDB.SetConnMaxLifetime(time.Second * 5)
	log.Println("Mysql is Collection!!!")
}
