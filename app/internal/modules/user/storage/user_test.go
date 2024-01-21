package storage

import (
	"context"
	"errors"
	"projects/LDmitryLD/repository/app/internal/db/adapter"
	"projects/LDmitryLD/repository/app/internal/db/adapter/mocks"
	"projects/LDmitryLD/repository/app/internal/models"
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
)

func TestUserStorage_Create(t *testing.T) {
	adapterer := mocks.NewSQLAdapterer(t)
	ctx := context.Background()
	testUser := models.UserDTO{
		ID:        1,
		FirstName: "name",
	}

	adapterer.On("Create", ctx, &testUser).Return(nil)

	userStorage := NewUserStorage(adapterer)

	err := userStorage.Create(ctx, testUser)

	assert.Nil(t, err)

}

func TestUserStorage_GetByID_NotFound(t *testing.T) {
	adapterer := mocks.NewSQLAdapterer(t)
	ID := 1
	ctx := context.Background()
	testDTO := models.User{}
	var testList []models.User
	testCond := adapter.Condition{
		Equal: map[string]interface{}{
			"id": ID,
		},
	}

	adapterer.On("List", ctx, &testList, &testDTO, testCond).Return(nil)

	userStorage := NewUserStorage(adapterer)

	_, err := userStorage.GetByID(ctx, ID)

	assert.NotNil(t, err)
}

func TestUserStorage_GetByID(t *testing.T) {
	adapterer := mocks.NewSQLAdapterer(t)
	ID := 1
	ctx := context.Background()
	testDTO := models.User{}
	var testList []models.User
	testCond := adapter.Condition{
		Equal: map[string]interface{}{
			"id": ID,
		},
	}

	adapterer.On("List", ctx, &testList, &testDTO, testCond).Return(errors.New("user storage: GetById not found user with ID 1"))

	userStorage := NewUserStorage(adapterer)

	_, err := userStorage.GetByID(ctx, ID)

	assert.NotNil(t, err)
}

func TestUserStorage_Update(t *testing.T) {
	adapterer := mocks.NewSQLAdapterer(t)

	ctx := context.Background()
	testUser := models.UserDTO{
		ID:        1,
		FirstName: "name",
	}
	testCond := adapter.Condition{
		Equal: sq.Eq{
			"id": testUser.GetID(),
		},
	}

	adapterer.On("Update", ctx, &testUser, testCond).Return(nil)

	userStorage := NewUserStorage(adapterer)

	err := userStorage.Update(ctx, testUser)

	assert.Nil(t, err)
}

func TestUserStorage_Delete(t *testing.T) {
	adapterer := mocks.NewSQLAdapterer(t)

	ctx := context.Background()
	tableName := "users"
	id := 1

	adapterer.On("Delete", ctx, tableName, id).Return(nil)

	userStorage := NewUserStorage(adapterer)

	err := userStorage.Delete(ctx, tableName, id)

	assert.Nil(t, err)
}

func TestUserStorage_List(t *testing.T) {
	adapterer := mocks.NewSQLAdapterer(t)

	testCond := adapter.Condition{
		Equal: sq.Eq{
			"id": 1,
		},
	}
	ctx := context.Background()
	var testDTO models.User
	var testList []models.User

	adapterer.On("List", ctx, &testList, &testDTO, testCond).Return(nil)

	userStorage := NewUserStorage(adapterer)

	_, err := userStorage.List(ctx, testCond)

	assert.Nil(t, err)
}
