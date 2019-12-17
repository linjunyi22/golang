package controllers

import (
	"fmt"
	"gin-gorm/models"
	"gin-gorm/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func insertTest() {
	models.OrmHandle(func(db *gorm.DB) error {
		p := make(chan int)
		for i := 0; i < 100; i++ {
			tmp := strconv.Itoa(i)
			u := models.User{
				Name: "test" + tmp,
			}
			go func(u models.User, i int) {
				db.Create(&u)
				p <- i
			}(u, i)
		}
		for m := 0; m < 100; m++ {
			<-p
		}
		return nil
	})
}

func getTest() []models.User {
	var u []models.User
	models.OrmHandle(func(db *gorm.DB) error {
		fmt.Println("db", db)
		err := db.Find(&u).Error
		return err
	})
	return u
}

func GetData(ctx *gin.Context) {
	utils.GetDBConf()
	res := getTest()
	ctx.JSON(http.StatusOK, res)
	return
}

func InsertData(ctx *gin.Context) {
	insertTest()
	ctx.JSON(http.StatusOK, gin.H{
		"insert": "success",
	})
	return
}
