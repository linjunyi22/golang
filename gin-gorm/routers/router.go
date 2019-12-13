package routers

import (
	"gin-gorm/controllers"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine = gin.Default()

func init() {
	rg := Router.Group("/v1/api")
	rg.GET("/data", controllers.GetData)
	rg.GET("/data/insert", controllers.InsertData)
}
