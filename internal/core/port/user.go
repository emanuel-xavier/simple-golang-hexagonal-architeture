package port

import (
	"github.com/emanuel-xavier/hexagonal-architerure/internal/core/domain"
)

// Primary ports

type UserUseCase interface {
	CreateUser(username string, id string) (*domain.User, error)
	GetUserById(id string) (*domain.User, error)
	GetAll() ([]domain.User, error)
}

// Secundary ports

type UserRepository interface {
	Save(user domain.User) error
	GetUserById(id string) (*domain.User, error)
	GetAll() ([]domain.User, error)
}
