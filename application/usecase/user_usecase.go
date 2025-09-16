package usecase

import (
	"nofu-be/domain"
	"nofu-be/domain/user"
)

type UserUseCaseImpl struct {
	repo domain.UserRepository
}

func NewUserUseCase(repo domain.UserRepository) domain.UserUseCase {
	return &UserUseCaseImpl{repo: repo}
}

func (uc *UserUseCaseImpl) GetAllUsers() ([]user.User, error) {
	return uc.repo.ListAll()
}

// CreateUser menyimpan user baru ke database
func (uc *UserUseCaseImpl) CreateUser(u user.User) error { return uc.repo.InsertUser(u) }
