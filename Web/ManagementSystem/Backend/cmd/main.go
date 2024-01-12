package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"

	cfg "course.project/management_system/config"
)

func register(r *gin.Engine) {

}

func initDB() {

}

func initCfg() {
	err := cfg.InitConfig()
	if err != nil {
		panic("init config error: " + err.Error())
	}
}

func main() {
	initCfg()
	initDB()

	r := gin.Default()
	register(r)
	r.Run()
}
