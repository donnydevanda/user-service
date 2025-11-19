package dto

import "github.com/google/uuid"

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	UUID        uuid.UUID `json:"uuid"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Role        string    `json:"role"`
	PhoneNumber string    `json:"phone_number"`
}

type LoginResponse struct {
	User  UserResponse `json:"user"`
	Token string       `json:"token"`
}

type RegisterRequest struct {
	Name        string `json:"name" binding:"required"`
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	ConfirmPass string `json:"confirmPassword" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	RoleID      uint
}

type RegisterResponse struct {
	User UserResponse `json:"user"`
}

type UpdateReqeuest struct {
	Name        string `json:"name" binding:"required"`
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password,omitempty"`
	ConfirmPass string `json:"confirmPassword,omitempty"`
	Email       string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	RoleID      uint
}
