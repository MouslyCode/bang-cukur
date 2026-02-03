package userModel

type CreateRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	RoleID   uint   `json:"role_id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UpdateRequest struct {
	Name   string `json:"name" binding:"required"`
	Email  string `json:"email" binding:"required,email"`
	RoleID uint   `json:"role_id" binding:"required"`
}
