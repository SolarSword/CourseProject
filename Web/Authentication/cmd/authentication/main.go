package main

import (
	cfg "course.project/authentication/config"
	DB "course.project/authentication/internal/common/db"
)

func initDB() {
	err := DB.Db.InitDB()
	if err != nil {
		panic("init database error: " + err.Error())
	}
}

func initCfg() {
	err := cfg.InitConfig()
	if err != nil {
		panic("init config error: " + err.Error())
	}
}

func main() {

}
