package service

import (
	"time"

	"github.com/Skapar/simple-rest/config"
	"github.com/Skapar/simple-rest/internal/models/dto"
	"github.com/Skapar/simple-rest/internal/models/entities"
	"github.com/Skapar/simple-rest/internal/repository"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	repo   repository.AuthRepository
	config *config.Config
}

func NewAuthService(repo repository.AuthRepository, config *config.Config) AuthService {
	return &AuthServiceImpl{repo: repo, config: config}
}

func (s *AuthServiceImpl) RegisterUser(user *entities.User) (string, string, error) {
	existingUser, err := s.repo.GetUserByEmail(user.Email)
	if err != nil {
		return "", "", errors.Wrap(err, "failed to check if user exists")
	}
	if existingUser != nil {
		return "", "", errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", "", errors.Wrap(err, "failed to hash password")
	}
	user.Password = string(hashedPassword)

	createdUser, err := s.repo.CreateUser(user)
	if err != nil {
		return "", "", errors.Wrap(err, "failed to create user")
	}

	accessToken, err := s.generateJWT(createdUser.ID)
	if err != nil {
		return "", "", errors.Wrap(err, "failed to generate access token")
	}

	refreshToken, err := s.generateRefreshToken(createdUser.ID)
	if err != nil {
		return "", "", errors.Wrap(err, "failed to generate refresh token")
	}

	createdUser.RefreshToken = refreshToken
	if _, err := s.repo.UpdateUser(createdUser); err != nil {
		return "", "", errors.Wrap(err, "failed to update user with refresh token")
	}

	return accessToken, refreshToken, nil
}

func (s *AuthServiceImpl) generateJWT(id int64) (string, error) {
	claims := dto.Claims{
		UserID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)), // 1 hour expiration
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.config.JWTSecret))
	if err != nil {
		return "", errors.New("error creating JWT")
	}

	return tokenString, nil
}

func (s *AuthServiceImpl) generateRefreshToken(id int64) (string, error) {
	claims := dto.Claims{
		UserID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)), // 1 week expiration
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.config.JWTSecret))
	if err != nil {
		return "", errors.New("error creating refresh token")
	}

	return tokenString, nil
}
