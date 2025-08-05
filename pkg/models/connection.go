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
		panic(err)
	}
	var connLifetime time.Duration = time.Duration(conf.Config.MYSQL_CONN_MAX_LIFETIME) * time.Second
	db.SetConnMaxLifetime(connLifetime)
	db.SetMaxOpenConns(conf.Config.MYSQL_MAX_OPEN_CONNS)
	db.SetMaxIdleConns(conf.Config.MYSQL_MAX_IDLE_CONNS)
}

func getDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", config.Config.MYSQL_USERNAME, config.Config.MYSQL_PASSWORD, config.Config.MYSQL_HOST, config.Config.MYSQL_DATABASE)
}
