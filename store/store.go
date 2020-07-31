package store

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Store is a struct that stores data to work with storage.
type Store struct {
	config          *Config
	db              *sql.DB
	pointRepository *PointRepository
}

// New creates new store.
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Open opens connection with database.
func (store *Store) Open() error {
	db, err := sql.Open("mysql", store.config.DatabaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	store.db = db

	return nil
}

// Ping checks connection to database.
func (store *Store) Ping() error {
	return store.db.Ping()
}

// Close closes connection with database.
func (store *Store) Close() {
	store.db.Close()
}

// Point entrypoint to work with point repository.
func (store *Store) Point() *PointRepository {
	if store.pointRepository != nil {
		return store.pointRepository
	}

	return &PointRepository{
		store: store,
	}
}
