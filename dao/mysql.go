package dao

import (
	"Todogo/setting"
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

// 初始化数据库
func InitMySQL(conf *setting.MySQLConfig) (err error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.User, conf.Password, conf.Host, conf.Port, conf.DB)
	Db, err = sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("connect server failed, err:%v\n", err)
		return err
	}
	Db.SetMaxOpenConns(200)
	Db.SetMaxIdleConns(10)

	initSql := `create table if not exists todo
				(
					id varchar(64) not null
						primary key,
					title varchar(64) null,
					status tinyint(1) null
				);`
	_, _ = Db.Exec(initSql)
	return Db.Ping()
}
