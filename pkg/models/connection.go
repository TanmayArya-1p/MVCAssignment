package models

import (
	"database/sql"
	"fmt"
	"time"

	"inorder/pkg/config"
	conf "inorder/pkg/config"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

const MYSQL_MAX_BIG_INT string = "18446744073709551610"

func init() {
	var dsn string = getDSN()
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Error opening database connection:", err)
		panic(err)
	}
	var connLifetime time.Duration = time.Duration(conf.Config.MySQL.CONN_MAX_LIFETIME) * time.Second
	db.SetConnMaxLifetime(connLifetime)
	db.SetMaxOpenConns(conf.Config.MySQL.MAX_OPEN_CONNS)
	db.SetMaxIdleConns(conf.Config.MySQL.MAX_IDLE_CONNS)
}

func getDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", config.Config.MySQL.USERNAME, config.Config.MySQL.PASSWORD, config.Config.MySQL.HOST, config.Config.MySQL.DATABASE)
}
