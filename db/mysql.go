package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func MysqlDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@/bltrbot?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	return db
}
