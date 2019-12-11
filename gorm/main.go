package main

import (
	"gorm/model"
	"github.com/jinzhu/gorm"
)

func insertTest() {
	u := model.User{
		Name: "test",
	}

	model.OrmHandle(func(db *gorm.DB) error {
		db.Create(&u)
		return nil
	})
}

func main() {
	insertTest()
}
