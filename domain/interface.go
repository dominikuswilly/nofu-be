package domain

import "nofu-be/domain/user"

type UserRepository interface {
	ListAll() ([]user.User, error)
	InsertUser(u user.User) error
}

type UserUseCase interface {
	GetAllUsers() ([]user.User, error)
	CreateUser(u user.User) error
}
