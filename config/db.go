package config

import "fmt"

type DBConfig struct {
	Type string
	Host string
	Name string
	UserName string
	Password string
	Port int
	PoolConns int
}

func (dbConfig *DBConfig) GetSettingStr() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		dbConfig.UserName,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
	)
}