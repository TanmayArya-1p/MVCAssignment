package config

type config struct {
	MySQL struct {
		USERNAME          string `yaml:"USERNAME"`
		PASSWORD          string `yaml:"PASSWORD"`
		HOST              string `yaml:"HOST"`
		DATABASE          string `yaml:"DATABASE"`
		CONN_MAX_LIFETIME int    `yaml:"CONN_MAX_LIFETIME"`
		MAX_OPEN_CONNS    int    `yaml:"MAX_OPEN_CONNS"`
		MAX_IDLE_CONNS    int    `yaml:"MAX_IDLE_CONNS"`
	} `yaml:"MySQL"`
	InOrder struct {
		PORT                 string `yaml:"PORT"`
		REFRESH_TOKEN_EXPIRE int    `yaml:"REFRESH_TOKEN_EXPIRE"`
		JTI_CLEANUP_INTERVAL int    `yaml:"JTI_CLEANUP_INTERVAL"`
	} `yaml:"InOrder"`
}
