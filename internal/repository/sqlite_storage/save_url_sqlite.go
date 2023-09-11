package sqlite_storage

import (
	"database/sql"
	"errors"
	"fmt"

	// "github.com/mattn/go-sqlite3"
	_ "github.com/mattn/go-sqlite3"
)

type SqliteStorage struct {
	db *sql.DB
}

func NewSqliteStorage(db *sql.DB) *SqliteStorage {
	return &SqliteStorage{db: db}
}

var (
	ErrURLNotFound = errors.New("url not found")
	ErrURLExists   = errors.New("url exists")
)

func (s *SqliteStorage) SaveShortUrl(shortURL, fullURL string) error {
	// // Вставляем данные в таблицу
	// const op = "storage.sqlite.SaveURL"
	fmt.Println(shortURL, fullURL)
	// stmt, err := s.db.Prepare("INSERT INTO urls(alias, fullurl) VALUES(?, ?)")
	// if err != nil {
	// 	return fmt.Errorf("%s: %w", op, err)
	// }
	// fmt.Println(shortURL, fullURL)
	// _, err = stmt.Exec(shortURL, fullURL)
	// if err != nil {
	// 	if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
	// 		return fmt.Errorf("%s: %w", op, ErrURLExists)
	// 	}

	// 	return fmt.Errorf("%s: %w", op, err)
	// }
	// return nil
	_, err := s.db.Exec("INSERT INTO urls (alias, fullurl) VALUES (?, ?)",
	shortURL, fullURL)
	if err != nil {
		return fmt.Errorf("error while inserting in databse: %w", err)
	}
	return nil
}

func (r *SqliteStorage) GetFullUrl(shortURL string) (string, error) {
	const op = "repository.save_url_sqlite.GetFullURL"

	stmt, err := r.db.Prepare("SELECT fullurl FROM urls WHERE alias = ?")
	if err != nil {
		return "", fmt.Errorf("%s: prepare statement: %w", op, err)
	}

	var resURL string

	err = stmt.QueryRow(shortURL).Scan(&resURL)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", ErrURLNotFound
		}

		return "", fmt.Errorf("%s: execute statement: %w", op, err)
	}

	return resURL, nil
}
