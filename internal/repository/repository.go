package repository

import (
	"github.com/Skapar/simple-rest/internal/models/entities"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &AuthRepositoryImpl{db: db}
}

func (r *AuthRepositoryImpl) CreateUser(user *entities.User) (*entities.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, errors.Wrap(err, "failed to create user")
	}
	return user, nil
}

func (r *AuthRepositoryImpl) GetUserByEmail(email string) (*entities.User, error) {
	var user entities.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user does not exist")
		}
		return nil, errors.Wrap(err, "failed to find user by email")
	}
	return &user, nil
}

func (r *AuthRepositoryImpl) GetUserById(userID int64) (*entities.User, error) {
	var user entities.User
	if err := r.db.Where("id = ?", userID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user does not exist")
		}
		return nil, errors.Wrap(err, "failed to find user by id")
	}
	return &user, nil
}

func (r *AuthRepositoryImpl) UpdateUser(user *entities.User) (*entities.User, error) {
	if err := r.db.Save(user).Error; err != nil {
		return nil, errors.Wrap(err, "failed to update user")
	}
	return user, nil
}

func (r *AuthRepositoryImpl) DeleteUser(user *entities.User) (*entities.User, error) {
	if err := r.db.Unscoped().Delete(user).Error; err != nil {
		return nil, errors.Wrap(err, "failed to delete user")
	}
	return user, nil
}

func (r *AuthRepositoryImpl) SoftDeleteUser(user *entities.User) (*entities.User, error) {
	if err := r.db.Delete(user).Error; err != nil {
		return nil, errors.Wrap(err, "failed to soft delete user")
	}
	return user, nil
}
