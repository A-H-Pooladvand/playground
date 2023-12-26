package mysql

import "po/configs"

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func NewConfig(
	host string,
	port string,
	username string,
	password string,
	database string,
) *Config {
	return &Config{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		Database: database,
	}
}

func NewEnvConfig() Config {
	c := configs.NewMysql()

	return Config{
		Host:     c.Host,
		Port:     c.Port,
		Username: c.Username,
		Password: c.Password,
		Database: c.Database,
	}
}