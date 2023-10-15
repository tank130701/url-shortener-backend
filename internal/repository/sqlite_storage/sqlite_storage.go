package sqlite_storage

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)


func NewSqliteDB(storagePath string)(*sql.DB, error){
	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()

	// Создаем таблицу, если она еще не существует
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS urls(
			id SERIAL,
			alias TEXT NOT NULL UNIQUE,
			fullurl TEXT NOT NULL);
		CREATE INDEX IF NOT EXISTS idx_alias ON urls(alias);
	`)
	if err != nil {
		return nil, fmt.Errorf("error creating database: %w", err)
	}
	return db, nil
}