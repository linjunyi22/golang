package controllers

import (
	"fmt"
	"gin-gorm/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func insertTest() {
	u := models.User{
		Name: "test",
	}

	models.OrmHandle(func(db *gorm.DB) error {
		db.Create(&u)
		return nil
	})
}

func getTest() []models.User {
	var u []models.User
	models.OrmHandle(func(db *gorm.DB) error {
		err := db.Find(&u).Error
		return err
	})
	return u
}

func GetData(ctx *gin.Context) {
	res := getTest()
	fmt.Println("res", res)
	ctx.JSON(http.StatusOK, res)
	return
}

func InsertData(ctx *gin.Context) {
	insertTest()
	ctx.JSON(http.StatusOK, gin.H{
		"hello": "world",
	})
	return
}
