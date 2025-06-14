package domain

type User struct {
	ID   string
	Name string
	Email string
}

type UserRepository interface {
	GetAll() ([]User, error)
	GetByID(id string) (User, error)
	Create(user User) error
	Delete(id string) error
}
