package service

import (
	"github.com/Skapar/simple-rest/internal/models/entities"
)

type AuthService interface {
	RegisterUser(user *entities.User) (string, string, error)
}

type UserService interface {
	GetUserById(userID int64) (*entities.User, error)
	UpdateUser(userID int64, user *entities.User) error
	DeleteUser(userID int64) (*entities.User, error)
	SoftDeleteUser(userID int64) (*entities.User, error)
}
