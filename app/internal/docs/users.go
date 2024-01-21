package docs

import "projects/LDmitryLD/repository/app/internal/modules/user/controller"

// swagger:route POST /api/users/create repo CreateRequest
// Создание пользователя в таблице.
// responses:
// 	200: CreateResponse

// swagger:parameters CreateRequest
type CreateRequest struct {
	// in:body
	Body controller.CreateRequest
}

// swagger:response CreateResponse
type CreateResponse struct {
	// in:body
	Body controller.CreateResponse
}

// swagger:route GET /api/users/{id} repo GetByIDRequest
// Получение пользователя по ID.
// responses:
// 	200: GetByIDResponse

// swagger:parameters GetByIDRequest
type GetByIDParameters struct {
	// in:path
	// required:true
	ID int `json:"id"`
}

// swagger:response GetByIDResponse
type GetByIDResponse struct {
	// in:body
	Body controller.GetByIDResponse
}

// swagger:route POST /api/users/update repo UpdateRequest
// Изменение пользователя.
// responses:
// 	200: UpdateResponse

// swagger:parameters UpdateRequest
type UpdateRequest struct {
	// in:body
	Body controller.UpdateRequest
}

// swagger:response UpdateResponse
type UpdateResponse struct {
	// in:body
	Body controller.UpdateResponse
}

// swagger:route POST /api/users/delete repo DeleteRequest
// Удаление пользователя по ID.
// responses:
// 	200: DeleteResponse

// swagger:parameters DeleteRequest
type DeleteRequest struct {
	// in:body
	Body controller.DeleteRequest
}

// swagger:response DeleteResponse
type DeleteResponse struct {
	// in:body
	Body controller.DeleteResponse
}

// swagger:route POST /api/users/list repo ListRequest
// Получение списка пользователей.
// responses:
// 	200: ListResponse

// swagger:parameters ListRequest
type ListRequest struct {
	// in:body
	Body controller.ListRequest
}

// swagger:response ListResponse
type ListResponse struct {
	// in:body
	Body controller.ListResponse
}
