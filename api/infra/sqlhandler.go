package infra

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type SqlHandler struct {
	Conn *gorm.DB
}

func NewSqlHandler() *SqlHandler {
	dsn := "root:hogehoge@tcp(db:3306)/dev?charset=utf8mb4&parseTime=True&loc=Local"
	conn, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err.Error)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}
