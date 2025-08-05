package config

import "time"

type config struct {
	MYSQL_USERNAME               string        `yaml:"MYSQL_USERNAME"`
	MYSQL_PASSWORD               string        `yaml:"MYSQL_PASSWORD"`
	MYSQL_HOST                   string        `yaml:"MYSQL_HOST"`
	MYSQL_DATABASE               string        `yaml:"MYSQL_DATABASE"`
	MYSQL_CONN_MAX_LIFETIME      int           `yaml:"MYSQL_CONN_MAX_LIFETIME"`
	MYSQL_MAX_OPEN_CONNS         int           `yaml:"MYSQL_MAX_OPEN_CONNS"`
	MYSQL_MAX_IDLE_CONNS         int           `yaml:"MYSQL_MAX_IDLE_CONNS"`
	INORDER_PORT                 string        `yaml:"PORT"`
	INORDER_REFRESH_TOKEN_EXPIRE int           `yaml:"REFRESH_TOKEN_EXPIRE"`
	INORDER_JTI_CLEANUP_INTERVAL time.Duration `yaml:"JTI_CLEANUP_INTERVAL"`
}
