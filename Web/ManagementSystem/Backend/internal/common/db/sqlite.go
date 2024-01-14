package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"

	cfg "course.project/management_system/config"
)

type SQLite struct {
	db *sql.DB
}

var Db SQLite

func (s *SQLite) InitDB() error {
	db, err := sql.Open(cfg.Cfg.Database.Driver, cfg.Cfg.Database.Path)
	if err != nil {
		log.Fatalf("get error: %v when opening sqlite: %v \n", err, cfg.Cfg.Database)
		return err
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("get error: %v when connecting to sqlite: %v \n", err, cfg.Cfg.Database)
		return err
	}
	s.db = db
	return nil
}

func (s *SQLite) CloseDB() {
	s.db.Close()
}
