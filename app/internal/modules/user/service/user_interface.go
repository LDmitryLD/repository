package service

import (
	"context"
	"projects/LDmitryLD/repository/app/internal/models"
)

type Userer interface {
	Create(ctx context.Context, in CreateIn) CreateOut
	GetByID(ctx context.Context, in GetByIDIn) GetByIDOut
	Update(ctx context.Context, in UpdateIn) UpdateOut
	Delete(ctx context.Context, in DeleteIn) DeleteOut
	List(ctx context.Context, in ListIn) ListOut
}

type CreateIn struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Address   string `json:"address"`
}

type CreateOut struct {
	Error error `json:"error"`
}

type GetByIDIn struct {
	UserID int `json:"user_id"`
}

type GetByIDOut struct {
	User  models.User `json:"user"`
	Error error       `json:"error"`
}

type UpdateIn struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Address   string `json:"address"`
}

type UpdateOut struct {
	Error error `json:"error"`
}

type DeleteIn struct {
	TableName string `json:"table_name"`
	UserID    int    `json:"user_id"`
}

type DeleteOut struct {
	Error error `json:"error"`
}

type ListIn struct {
	Limit  int64
	Offset int64
}

type ListOut struct {
	// Users []models.UserDTO `json:"users"`
	Users []models.User `json:"users"`
	Error error         `json:"error"`
}
