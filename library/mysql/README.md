## sql

- 第三方库:
	github.com/jmoiron/sqlx
	
- 第三方引擎:
	github.com/go-sql-driver/mysql
	
- 使用:
	- 调用NewMysql 传入配置文件
        address :ip 
        idleConn : 用于设置闲置的连接数。
        openConn : 用于设置最大打开的连接数，默认值为0表示不限制。
	- 使用deps包下的Mysql获取连接
