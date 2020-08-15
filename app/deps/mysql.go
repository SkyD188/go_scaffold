package deps

import (
	"go_scaffold/library/mysql"

	"github.com/jmoiron/sqlx"
)

var (
	MysqlDB *sqlx.DB
)

func InitMysql() {
	MysqlDB = mysql.New(&AppConfig.Mysql)
}
