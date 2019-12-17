package models

import (
	"gin-gorm/utils"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

// 初始化 DB
func initDB() {
	var err error
	DB, err = gorm.Open("mysql", utils.GetDBConf())
	utils.CheckErr(err)

	DB.AutoMigrate(&User{})
	// 设置空闲连接池中的最大连接数
	DB.DB().SetMaxIdleConns(10)
	// 设置到数据库的最大打开连接数
	DB.DB().SetMaxOpenConns(100)
}

// 回调处理
func OrmHandle(callback func(*gorm.DB) error) {
	tx := DB.Begin()
	err := callback(tx)
	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
}

func init() {
	initDB()
}
