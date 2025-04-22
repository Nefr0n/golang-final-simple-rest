package repository

import (
	"github.com/Skapar/simple-rest/internal/models/entities"
)

type AuthRepository interface {
	CreateUser(user *entities.User) (*entities.User, error)
	GetUserByEmail(email string) (*entities.User, error)
	UpdateUser(user *entities.User) (*entities.User, error)
	DeleteUser(user *entities.User) (*entities.User, error)
	SoftDeleteUser(user *entities.User) (*entities.User, error)
	GetUserById(id int64) (*entities.User, error)
}
