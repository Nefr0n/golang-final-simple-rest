package api

import (
	"net/http"
	"strconv"

	"github.com/Skapar/simple-rest/internal/models"
	"github.com/Skapar/simple-rest/internal/models/dto"
	"github.com/Skapar/simple-rest/internal/models/entities"
	"github.com/Skapar/simple-rest/internal/service"
	"github.com/Skapar/simple-rest/internal/utils"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// GetUserHandler godoc
// @Summary Get a user by ID
// @Description Get a user by ID
// @Tags user
// @Accept  json
// @Produce  json
// @Param   user_id path int true "User ID"
// @Success 200 {object} models.Response[entities.User]
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/user/{user_id} [get]
func (uh *UserHandler) GetUserHandler(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "invalid user_id")
		return
	}

	user, err := uh.userService.GetUserById(userID)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.Response[entities.User]{
		Data:    *user,
		Message: "user retrieved successfully",
		Success: true,
	})
}

// UpdateUserHandler godoc
// @Summary Update a user by ID
// @Description Update a user by ID
// @Tags user
// @Accept  json
// @Produce  json
// @Param   user_id path int true "User ID"
// @Param   user body dto.UpdateUserDTO true "User"
// @Success 200 {object} models.Response[entities.User]
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/user/{user_id} [put]
func (uh *UserHandler) UpdateUserHandler(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "invalid user_id")
		return
	}

	var userDTO dto.UpdateUserDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	user := entities.User{
		Username: userDTO.Username,
		Email:    userDTO.Email,
		Password: userDTO.Password,
	}

	if err := uh.userService.UpdateUser(userID, &user); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.Response[entities.User]{
		Data:    user,
		Message: "user updated successfully",
		Success: true,
	})
}

// DeleteUserHandler godoc
// @Summary Delete a user by ID
// @Description Delete a user by ID
// @Tags user
// @Accept  json
// @Produce  json
// @Param   user_id path int true "User ID"
// @Success 200 {object} models.Response[entities.User]
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/user/{user_id} [delete]
func (uh *UserHandler) DeleteUserHandler(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "invalid user_id")
		return
	}

	deletedUser, err := uh.userService.DeleteUser(userID)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.Response[entities.User]{
		Data:    *deletedUser,
		Message: "user updated successfully",
		Success: true,
	})
}

// SoftDeleteUserHandler godoc
// @Summary Soft delete a user by ID
// @Description Soft delete a user by ID
// @Tags user
// @Accept  json
// @Produce  json
// @Param   user_id path int true "User ID"
// @Success 200 {object} models.Response[entities.User]
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/user/{user_id} [patch]
func (uh *UserHandler) SoftDeleteUserHandler(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "invalid user_id")
		return
	}

	softDeletedUser, err := uh.userService.SoftDeleteUser(userID)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.Response[entities.User]{
		Data:    *softDeletedUser,
		Message: "user soft deleted successfully",
		Success: true,
	})
}
