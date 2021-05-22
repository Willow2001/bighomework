package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)
var (
	db *sqlx.DB
)
func InitDB(dns string) (err error) {

	db, err = sqlx.Open("mysql", dns)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return err
	} else{
		fmt.Printf("成功连接\n")
	}
	db.SetMaxOpenConns(100)//最大连接
	db.SetMaxIdleConns(16)//空闲

	return
}
