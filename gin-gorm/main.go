package main

import (
	"gin-gorm/routers"
)

func ginRun() {
	r := routers.Router
	r.Run()
}

func main() {
	ginRun()
}
