package repository

import (
	"GoLang-Notes/clean-architecture/internal/domain"
	"errors"
)

type InMemoryUserRepo struct {
	data map[string]domain.User
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{
		data: make(map[string]domain.User),
	}
}

func (r *InMemoryUserRepo) GetAll() ([]domain.User, error) {
	users := make([]domain.User, 0, len(r.data))
	for _, u := range r.data {
		users = append(users, u)
	}
	return users, nil
}

func (r *InMemoryUserRepo) GetByID(id string) (domain.User, error) {
	u, ok := r.data[id]
	if !ok {
		return domain.User{}, errors.New("user not found")
	}
	return u, nil
}

func (r *InMemoryUserRepo) Create(user domain.User) error {
	r.data[user.ID] = user
	return nil
}

func (r *InMemoryUserRepo) Delete(id string) error {
	delete(r.data, id)
	return nil
}
