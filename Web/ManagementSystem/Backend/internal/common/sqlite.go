package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"

	cfg "course.project/management_system/config"
)

// DB close TBD
func InitDB() (*sql.DB, error) {
	db, err := sql.Open(cfg.Cfg.Database.Driver, cfg.Cfg.Database.Path)
	if err != nil {
		log.Fatalf("get error: %v when loading open sqlite: %v \n", err, cfg.Cfg.Database)
		return nil, err
	}
	return db, nil
}
