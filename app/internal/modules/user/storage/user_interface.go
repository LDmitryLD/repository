package storage

import (
	"context"
	"projects/LDmitryLD/repository/app/internal/db/adapter"
	"projects/LDmitryLD/repository/app/internal/models"
)

//go:generate go run github.com/vektra/mockery/v2@v2.35.4 --name=UserRepository
type UserRepository interface {
	Create(ctx context.Context, user models.UserDTO) error
	GetByID(ctx context.Context, id int) (models.User, error)
	Update(ctx context.Context, user models.UserDTO) error
	Delete(ctx context.Context, tableName string, id int) error
	List(ctx context.Context, c adapter.Condition) ([]models.User, error)
}

// 	List(ctx context.Context, c adapter.Condition) ([]models.UserDTO, error)
