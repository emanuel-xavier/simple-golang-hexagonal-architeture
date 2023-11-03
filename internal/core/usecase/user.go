package usecase

import (
	"github.com/emanuel-xavier/hexagonal-architerure/internal/core/domain"
	"github.com/emanuel-xavier/hexagonal-architerure/internal/core/port"
)

func NewUserUseCase(userRepo port.UserRepository) port.UserUseCase {
	return userUseCase{
		userRepo: userRepo,
	}
}

// Implements port.UserUseCase
type userUseCase struct {
	userRepo port.UserRepository
}

func (u userUseCase) CreateUser(username string, id string) (*domain.User, error) {
	user := domain.NewUser(username, id)
	err := u.userRepo.Save(user)

	return &user, err
}

func (u userUseCase) GetUserById(id string) (user *domain.User, err error) {
	return u.userRepo.GetUserById(id)
}

func (u userUseCase) GetAll() ([]domain.User, error) {
	return u.userRepo.GetAll()
}
