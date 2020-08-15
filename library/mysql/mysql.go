package mysql

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Mysql
type Conf struct {
	Address     string
	MaxOpenConn int
	MaxIdleConn int
}

func New(cfg *Conf) *sqlx.DB {
	mysqlDB, err := sqlx.Open("mysql", cfg.Address)
	if err != nil {
		panic(err)
	}
	err = mysqlDB.Ping()
	if err != nil {
		panic(err)
	}
	mysqlDB.SetMaxIdleConns(cfg.MaxIdleConn)
	mysqlDB.SetMaxOpenConns(cfg.MaxOpenConn)
	mysqlDB.SetConnMaxLifetime(time.Second * 5)
	log.Println("Mysql is Collection!!!")
	return mysqlDB
}
