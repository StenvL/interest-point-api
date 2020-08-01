package store

import (
	"database/sql"

	"github.com/StenvL/interest-points-api/models/domain"
)

// UserRepository repository for working with users.
type UserRepository struct {
	store *Store
}

// GetByLogin returns user by login.
func (r *UserRepository) GetByLogin(login string) (*domain.User, error) {
	user := &domain.User{}
	err := r.store.db.QueryRow(
		"SELECT id, login, encrypted_password FROM user WHERE login = ?",
		login,
	).Scan(&user.ID, &user.Login, &user.EncryptedPassword)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}
