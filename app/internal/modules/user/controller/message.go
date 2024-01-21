package controller

import "projects/LDmitryLD/repository/app/internal/models"

const (
	LogErrDecodeReq = "ошибка при декодировании запроса"
)

type CreateRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Address   string `json:"address"`
}

type CreateResponse struct {
	Success bool  `json:"success"`
	Error   error `json:"error"`
}

// type GetByIDRequest struct {
// 	ID int `json:"id"`
// }

type GetByIDResponse struct {
	Success bool        `json:"success"`
	Error   error       `json:"error"`
	User    models.User `json:"user"`
}

type UpdateRequest struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Address   string `json:"address"`
}

type UpdateResponse struct {
	Success bool  `json:"success"`
	Error   error `json:"error"`
}

type DeleteRequest struct {
	TableName string `json:"table_name"`
	ID        int    `json:"id"`
}

type DeleteResponse struct {
	Success bool  `json:"success"`
	Error   error `json:"error"`
}

type ListRequest struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

type ListResponse struct {
	Success bool  `json:"success"`
	Error   error `json:"error"`
	// Users   []models.UserDTO `json:"users"`
	Users []models.User `json:"users"`
}
