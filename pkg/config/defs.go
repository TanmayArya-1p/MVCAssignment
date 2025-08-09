package config

type MySQLConfig struct {
	USERNAME          string `yaml:"USERNAME"`
	PASSWORD          string `yaml:"PASSWORD"`
	HOST              string `yaml:"HOST"`
	DATABASE          string `yaml:"DATABASE"`
	CONN_MAX_LIFETIME int    `yaml:"CONN_MAX_LIFETIME"`
	MAX_OPEN_CONNS    int    `yaml:"MAX_OPEN_CONNS"`
	MAX_IDLE_CONNS    int    `yaml:"MAX_IDLE_CONNS"`
}

type InOrderConfig struct {
	PORT                 string `yaml:"PORT"`
	REFRESH_TOKEN_EXPIRE int    `yaml:"REFRESH_TOKEN_EXPIRE"`
	JTI_CLEANUP_INTERVAL int    `yaml:"JTI_CLEANUP_INTERVAL"`
	JWT_SECRET           string `yaml:"JWT_SECRET"`
	AUTH_TOKEN_EXPIRY    int    `yaml:"AUTH_TOKEN_EXPIRY"`
	REFRESH_TOKEN_EXPIRY int    `yaml:"REFRESH_TOKEN_EXPIRY"`
	ITEM_IMAGE_DIRECTORY string `yaml:"ITEM_IMAGE_DIRECTORY"`
}

type config struct {
	MySQL   MySQLConfig   `yaml:"MySQL"`
	InOrder InOrderConfig `yaml:"InOrder"`
}
