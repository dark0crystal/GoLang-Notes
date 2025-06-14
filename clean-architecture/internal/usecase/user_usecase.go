package usecase

import "GoLang-Notes/clean-architecture/internal/domain"

type UserUseCase struct {
	repo domain.UserRepository
}

func NewUserUseCase(repo domain.UserRepository) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (uc *UserUseCase) GetAllUsers() ([]domain.User, error) {
	return uc.repo.GetAll()
}

func (uc *UserUseCase) CreateUser(user domain.User) error {
	return uc.repo.Create(user)
}

func (uc *UserUseCase) DeleteUser(id string) error {
	return uc.repo.Delete(id)
}
