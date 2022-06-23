package config

import (
	"fmt"
	"os"
)

var Mysql mysql

type mysql struct{}

func (mysql) UserName() string {
	return os.Getenv("MYSQL_USERNAME")
}

func (mysql) Password() string {
	return os.Getenv("MYSQL_PASSWORD")
}

func (mysql) Database() string {
	return os.Getenv("MYSQL_DATABASE")
}

func MysqlDSN() string {
	return fmt.Sprintf("%s:%s@/%s?parseTime=true", Mysql.UserName(), Mysql.Password(), Mysql.Database())
}
