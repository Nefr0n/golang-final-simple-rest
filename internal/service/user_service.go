package service

import (
	"fmt"
	"github.com/Skapar/simple-rest/internal/models/entities"
	"github.com/Skapar/simple-rest/internal/repository"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	repo repository.AuthRepository
}

func NewUserService(repo repository.AuthRepository) UserService {
	return &UserServiceImpl{repo: repo}
}

func (s *UserServiceImpl) GetUserById(userID int64) (*entities.User, error) {
	user, err := s.repo.GetUserById(userID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user")
	}
	return user, nil
}

func (s *UserServiceImpl) UpdateUser(userID int64, user *entities.User) error {
	existingUser, err := s.repo.GetUserById(userID)
	fmt.Println(existingUser)
	if err != nil {
		return errors.Wrap(err, "failed to get existing user")
	}
	if existingUser == nil {
		return errors.New("user does not exist")
	}

	if err := s.updateUsername(existingUser, user.Username); err != nil {
		return err
	}
	if err := s.updateEmail(existingUser, user.Email, userID); err != nil {
		return err
	}
	if err := s.updatePassword(existingUser, user.Password); err != nil {
		return err
	}

	if _, err := s.repo.UpdateUser(existingUser); err != nil {
		return errors.Wrap(err, "failed to update user")
	}
	return nil
}

func (s *UserServiceImpl) updateUsername(existingUser *entities.User, username string) error {
	if username != "" {
		existingUser.Username = username
	}
	return nil
}

func (s *UserServiceImpl) updateEmail(existingUser *entities.User, email string, userID int64) error {
	if email != "" {
		//otherUser, err := s.repo.GetUserByEmail(email)
		//if err != nil {
		//	return errors.Wrap(err, "failed to check if email is already in use")
		//}
		//if otherUser != nil && otherUser.ID != userID {
		//	return errors.New("email is already in use by another user")
		//}
		existingUser.Email = email
	}
	return nil
}

func (s *UserServiceImpl) updatePassword(existingUser *entities.User, password string) error {
	if password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return errors.Wrap(err, "failed to hash password")
		}
		existingUser.Password = string(hashedPassword)
	}
	return nil
}

func (s *UserServiceImpl) DeleteUser(userID int64) (*entities.User, error) {
	user := &entities.User{ID: userID}
	deletedUser, err := s.repo.DeleteUser(user)
	if err != nil {
		return nil, errors.Wrap(err, "failed to delete user")
	}
	return deletedUser, nil
}

func (s *UserServiceImpl) SoftDeleteUser(userID int64) (*entities.User, error) {
	user := &entities.User{ID: userID}
	softDeletedUser, err := s.repo.SoftDeleteUser(user)
	if err != nil {
		return nil, errors.Wrap(err, "failed to soft delete user")
	}
	return softDeletedUser, nil
}
