package utils

import (
	"fmt"

	"github.com/Unknwon/goconfig"
)

// 检查 error 直接panic
func CheckErr(e error) {
	if e != nil {
		panic(e.Error())
	}
}

// 获取 DB 配置
func GetDBConf() string {
	path := "./conf/app.ini"
	cfg, err := goconfig.LoadConfigFile(path)
	CheckErr(err)

	sec, err := cfg.GetSection("mysql")
	CheckErr(err)

	return fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True&loc=Local", sec["username"], sec["password"], sec["dbname"], sec["charset"])
}
