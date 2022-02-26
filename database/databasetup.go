package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/join-api/model"
)

func SetupModels() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@(localhost)/join?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&model.User{})
	return db
}
