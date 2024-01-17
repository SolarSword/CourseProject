package db

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	cfg "course.project/management_system/config"
)

type SQLite struct {
	db *gorm.DB
}

var Db SQLite

func (s *SQLite) InitDB() error {
	db, err := gorm.Open(sqlite.Open(cfg.Cfg.Database.Path), &gorm.Config{})
	//db, err := sql.Open(cfg.Cfg.Database.Driver, cfg.Cfg.Database.Path)
	if err != nil {
		log.Fatalf("get error: %v when opening sqlite: %v \n", err, cfg.Cfg.Database)
		return err
	}

	s.db = db
	return nil
}

func (s *SQLite) CloseDB() {
	if db, _ := s.db.DB(); db != nil {
		_ = db.Close()
	}
}

func (s *SQLite) GetDB() *gorm.DB {
	return s.db
}
