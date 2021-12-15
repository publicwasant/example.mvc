package controller

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DBConfig struct {
	Ip   string
	Port string
	User string
	Name string
	Conn *sql.DB
}

var DB_CONFIG DBConfig

func GetDatabaseConnection() DBConfig {
	var config DBConfig = DBConfig{`127.0.0.1`, `3306`, `root`, `t_dm`, nil}
	var src string = fmt.Sprintf(`%v@tcp(%v:%v)/%v`, config.User, config.Ip, config.Port, config.Name)

	if db, err := sql.Open("mysql", src); err == nil {
		config.Conn = db
	} else {
		fmt.Println(err)
	}

	return config
}
