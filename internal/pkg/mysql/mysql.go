package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func getMysqlDB() *sqlx.DB {
	return mysqlDB
}

type mysqlStruct struct {
	db *sqlx.DB
}

func MCurd() Iface {
	return &mysqlStruct{
		db: getMysqlDB(),
	}
}

func (m *mysqlStruct) Find()    {}
func (m *mysqlStruct) FindOne() {}
func (m *mysqlStruct) Insert()  {}
func (m *mysqlStruct) Update()  {}
func (m *mysqlStruct) Delete()  {}
