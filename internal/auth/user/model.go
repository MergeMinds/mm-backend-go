package user

import (
	"time"

	"github.com/google/uuid"
)

type Model struct {
	Id           uuid.UUID
	CreatedAt    time.Time
	FirstName    string
	LastName     string
	Username     string
	Email        string
	Role         string
	PasswordHash []byte
	PasswordSalt []byte
}

type OutModel struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
}

type CreateModel struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Role      string `json:"role" binding:"required"`
	Password  string `json:"password" binding:"password"`
}
