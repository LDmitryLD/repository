package service

import (
	"context"
	"projects/LDmitryLD/repository/app/internal/db/adapter"
	"projects/LDmitryLD/repository/app/internal/models"
	"projects/LDmitryLD/repository/app/internal/modules/user/storage/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserService_Create(t *testing.T) {
	userRepo := mocks.NewUserRepository(t)
	testIn := CreateIn{
		FirstName: "name",
		LastName:  "lastname",
		Username:  "username",
	}
	testDTO := models.UserDTO{
		FirstName: "name",
		LastName:  "lastname",
		Username:  "username",
	}

	userRepo.On("Create", context.Background(), testDTO).Return(nil)

	userService := NewUserService(userRepo)

	out := userService.Create(context.Background(), testIn)

	assert.Equal(t, nil, out.Error)
}

func TestUserService_GetByID(t *testing.T) {
	userRepo := mocks.NewUserRepository(t)
	testIn := GetByIDIn{
		UserID: 1,
	}
	testUser := models.User{
		LastName: "name",
	}
	expect := GetByIDOut{
		User:  testUser,
		Error: nil,
	}

	userRepo.On("GetByID", context.Background(), testIn.UserID).Return(testUser, nil)

	userService := NewUserService(userRepo)

	out := userService.GetByID(context.Background(), testIn)

	assert.Equal(t, expect, out)
}

func TestUserService_Update(t *testing.T) {
	userRepo := mocks.NewUserRepository(t)
	testIn := UpdateIn{
		ID:       1,
		LastName: "name",
	}
	testDTO := models.UserDTO{
		ID:       1,
		LastName: "name",
	}
	expect := UpdateOut{
		Error: nil,
	}

	userRepo.On("Update", context.Background(), testDTO).Return(nil)

	userService := NewUserService(userRepo)

	out := userService.Update(context.Background(), testIn)

	assert.Equal(t, expect, out)
}

func TestUserService_Delete(t *testing.T) {
	userRepo := mocks.NewUserRepository(t)
	testIn := DeleteIn{
		TableName: "users",
		UserID:    1,
	}
	expect := DeleteOut{
		Error: nil,
	}

	userRepo.On("Delete", context.Background(), testIn.TableName, testIn.UserID).Return(nil)

	userService := NewUserService(userRepo)

	out := userService.Delete(context.Background(), testIn)

	assert.Equal(t, expect, out)
}

func TestUserService_List(t *testing.T) {
	userRepo := mocks.NewUserRepository(t)
	testIn := ListIn{
		Limit:  3,
		Offset: 1,
	}
	testCond := adapter.Condition{
		LimitOffset: &adapter.LimitOffset{
			Limit:  testIn.Limit,
			Offset: testIn.Offset,
		},
	}
	testUsers := []models.User{{ID: 1, LastName: "name"}, {ID: 2, LastName: "name"}}
	expect := ListOut{
		Users: testUsers,
		Error: nil,
	}

	userRepo.On("List", context.Background(), testCond).Return(testUsers, nil)

	userService := NewUserService(userRepo)

	out := userService.List(context.Background(), testIn)

	assert.Equal(t, expect, out)
}
