package store

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Store struct {
	config *Config
	db     *sql.DB
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (store *Store) Open() error {
	db, err := sql.Open("mysql", store.config.DatabaseUrl)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	store.db = db

	return nil
}

func (store *Store) Close() {
	store.db.Close()
}
