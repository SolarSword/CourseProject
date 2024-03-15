package main

import (
	cfg "course.project/authentication/config"
	cache "course.project/authentication/internal/common/cache"
	DB "course.project/authentication/internal/common/db"
)

func initCfg() {
	err := cfg.InitConfig()
	if err != nil {
		panic("init config error: " + err.Error())
	}
}

func initCache() {
	err := cache.Ca.InitCache()
	if err != nil {
		panic("init cache error: " + err.Error())
	}
}

func initDB() {
	err := DB.Db.InitDB()
	if err != nil {
		panic("init database error: " + err.Error())
	}
}
