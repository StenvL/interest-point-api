package domain

// User domain model.
type User struct {
	ID                uint64
	Login             string
	EncryptedPassword string
}
