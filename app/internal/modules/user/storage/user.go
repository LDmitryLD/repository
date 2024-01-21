package storage

import (
	"context"
	"fmt"
	"projects/LDmitryLD/repository/app/internal/db/adapter"
	"projects/LDmitryLD/repository/app/internal/models"

	sq "github.com/Masterminds/squirrel"
)

type UserStorage struct {
	adapter adapter.SQLAdapterer
}

func NewUserStorage(sqlAdapter adapter.SQLAdapterer) *UserStorage {
	return &UserStorage{adapter: sqlAdapter}
}

func (s *UserStorage) Create(ctx context.Context, user models.UserDTO) error {
	return s.adapter.Create(ctx, &user)
}

func (s *UserStorage) GetByID(ctx context.Context, id int) (models.User, error) {
	var user models.User
	var list []models.User

	err := s.adapter.List(ctx, &list, &user, adapter.Condition{
		Equal: map[string]interface{}{
			"id": id,
		},
	})

	if err != nil {
		return models.User{}, err
	}
	if len(list) < 1 {
		return models.User{}, fmt.Errorf("user storage: GetById not found user with ID %d", id)
	}

	return list[0], nil
}

func (s *UserStorage) Update(ctx context.Context, user models.UserDTO) error {
	return s.adapter.Update(ctx, &user, adapter.Condition{
		Equal: sq.Eq{
			"id": user.GetID(),
		},
	})

}

func (s *UserStorage) Delete(ctx context.Context, tableName string, id int) error {
	return s.adapter.Delete(ctx, tableName, id)
}

func (s *UserStorage) List(ctx context.Context, c adapter.Condition) ([]models.User, error) {
	var user models.User
	var list []models.User

	err := s.adapter.List(ctx, &list, &user, c)

	return list, err
}
