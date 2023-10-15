package sqlite_storage

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/mattn/go-sqlite3"
	// _ "github.com/mattn/go-sqlite3"
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

// func (s *SqliteStorage) SaveShortUrl(urlToSave string, alias string) (int64, error) {
func (s *SqliteStorage) SaveShortUrl(urlToSave string, alias string)error {
	const op = "storage.sqlite.SaveURL"

	stmt, err := s.db.Prepare("INSERT INTO url(url, alias) VALUES(?, ?)")
	if err != nil {
		// return 0, fmt.Errorf("%s: %w", op, err)
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec(urlToSave, alias)
	if err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			// return 0, fmt.Errorf("%s: %w", op, ErrURLExists)
			return fmt.Errorf("%s: %w", op, ErrURLExists)
		}

		// return 0, fmt.Errorf("%s: %w", op, err)
		return fmt.Errorf("%s: %w", op, err)
	}

	// id, err := res.LastInsertId()
	// if err != nil {
	// 	return 0, fmt.Errorf("%s: failed to get last insert id: %w", op, err)
	// }

	// return id, nil
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
