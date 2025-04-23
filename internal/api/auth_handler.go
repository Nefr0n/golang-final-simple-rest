package api

import (
	"net/http"

	"github.com/Skapar/simple-rest/internal/models"
	"github.com/Skapar/simple-rest/internal/models/dto"
	"github.com/Skapar/simple-rest/internal/models/entities"
	"github.com/Skapar/simple-rest/internal/service"
	"github.com/Skapar/simple-rest/internal/utils"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// RegisterUserHandler godoc
// @Summary Register a new user
// @Description Register a new user with the input payload
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   user body dto.CreateUserDTO true "User"
// @Success 201 {object} models.Response[map[string]string]
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/signup [post]
func (ah *AuthHandler) RegisterUserHandler(c *gin.Context) {
	var userDTO dto.CreateUserDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	user := entities.User{
		Username: userDTO.Username,
		Email:    userDTO.Email,
		Password: userDTO.Password,
	}

	accessToken, refreshToken, err := ah.authService.RegisterUser(&user)
	if err != nil {
		handleRegisterUserError(c, err)
		return
	}

	c.JSON(http.StatusCreated, models.Response[interface{}]{
		Data:    gin.H{"access_token": accessToken, "refresh_token": refreshToken},
		Message: "user registered successfully",
		Success: true,
	})
}

func handleRegisterUserError(c *gin.Context, err error) {
	switch err.Error() {
	case "user already exists":
		utils.RespondWithError(c, http.StatusConflict, err.Error())
	default:
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
	}
}
