package service

import (
	"context"
	"projects/LDmitryLD/repository/app/internal/db/adapter"
	"projects/LDmitryLD/repository/app/internal/models"
	"projects/LDmitryLD/repository/app/internal/modules/user/storage"
)

type UserService struct {
	storage storage.UserRepository
}

func NewUserService(storage storage.UserRepository) *UserService {
	return &UserService{storage: storage}
}

func (u *UserService) Create(ctx context.Context, in CreateIn) CreateOut {
	dto := models.UserDTO{
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Username:  in.Username,
		Email:     in.Email,
		Address:   in.Address,
	}

	err := u.storage.Create(ctx, dto)

	return CreateOut{
		Error: err,
	}

}

func (u *UserService) GetByID(ctx context.Context, in GetByIDIn) GetByIDOut {
	user, err := u.storage.GetByID(ctx, in.UserID)
	// if err != nil {
	// 	return GetByIDOut{
	// 		Error: err,
	// 	}
	// }

	return GetByIDOut{
		User:  user,
		Error: err,
	}
}

// возможно стоить добавить в models стрктуру user без поля DeletedAt и оперировать ей
func (u *UserService) Update(ctx context.Context, in UpdateIn) UpdateOut {
	dto := models.UserDTO{
		ID:        in.ID,
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Username:  in.Username,
		Email:     in.Email,
		Address:   in.Address,
	}

	err := u.storage.Update(ctx, dto)

	return UpdateOut{
		Error: err,
	}
}

func (u *UserService) Delete(ctx context.Context, in DeleteIn) DeleteOut {
	err := u.storage.Delete(ctx, in.TableName, in.UserID)

	return DeleteOut{
		Error: err,
	}
}

func (u *UserService) List(ctx context.Context, in ListIn) ListOut {
	condition := adapter.Condition{
		LimitOffset: &adapter.LimitOffset{
			Limit:  in.Limit,
			Offset: in.Offset,
		},
	}

	users, err := u.storage.List(ctx, condition)

	return ListOut{
		Users: users,
		Error: err,
	}
}
