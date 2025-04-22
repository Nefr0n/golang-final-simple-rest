package dto

type User struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type CreateUserDTO struct {
	User
	Password string `json:"password" binding:"required"`
}

type UpdateUserDTO struct {
	User
	Password string `json:"password"`
}
