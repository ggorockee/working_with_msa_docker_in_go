package database

type ResponseType interface {
	ResponseUser | ResponseTinyUser
}

type Repository[T ResponseType] interface {
	GetByEmail(email string) error
	Create() error
	Serialize() T
}
