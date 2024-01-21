package storage

import (
	"context"
	"fmt"
	"projects/LDmitryLD/repository/app/internal/db/adapter"
	"projects/LDmitryLD/repository/app/internal/models"

	sq "github.com/Masterminds/squirrel"
)

type UserStorage struct {
	adapter *adapter.SQLAdapter
}

func NewUserStorage(sqlAdapter *adapter.SQLAdapter) *UserStorage {
	return &UserStorage{adapter: sqlAdapter}
}

func (s *UserStorage) Create(ctx context.Context, user models.UserDTO) error {
	return s.adapter.Create(ctx, &user)
}

// изменил UserDTO на User
func (s *UserStorage) GetByID(ctx context.Context, id int) (models.User, error) {
	// var dto models.UserDTO
	// var list []models.UserDTO

	var dto models.User
	var list []models.User

	err := s.adapter.List(ctx, &list, &dto, adapter.Condition{
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

// изменил UserDTO на User
func (s *UserStorage) List(ctx context.Context, c adapter.Condition) ([]models.User, error) {
	// var dto models.UserDTO
	// var list []models.UserDTO

	var dto models.User
	var list []models.User

	err := s.adapter.List(ctx, &list, &dto, c)
	if err != nil {
		return nil, err
	}

	// пройтись циклом по массиму и удалять специальной функцией пользователей которые помечены как удалённые?

	return list, nil
}
