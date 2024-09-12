package database

type Repository interface {
	GetByEmail(email string) error
	Create() error
}
