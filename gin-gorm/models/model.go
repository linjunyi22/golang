package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


type DB_CONFIG struct {
	username string
	password string
	dbName   string
	charset  string
}

func getDbConfig() string {
	var dbParams = DB_CONFIG{
		username: "root",
		password: "password",
		dbName:   "hello",
		charset:  "utf8",
	}
	return fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True&loc=Local", dbParams.username, dbParams.password, dbParams.dbName, dbParams.charset)
}

var DB *gorm.DB

// 初始化 DB
func initDB() {
	var err error
	DB, err = gorm.Open("mysql", getDbConfig())
	if err != nil {
		os.Exit(-1)
		panic(err)
	}

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
