package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"projects/LDmitryLD/repository/app/internal/infrastructure/responder"
	"projects/LDmitryLD/repository/app/internal/modules/user/service"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Userer interface {
	Craete(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
}

type User struct {
	service service.Userer
	responder.Responder
}

func NewUser(service service.Userer) Userer {
	return &User{
		service:   service,
		Responder: &responder.Respond{},
	}
}

func (u *User) Craete(w http.ResponseWriter, r *http.Request) {
	var req CreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(LogErrDecodeReq)
		u.ErrorBadRequest(w, err)
	}

	out := u.service.Create(r.Context(), service.CreateIn{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Username:  req.Username,
		Email:     req.Email,
		Address:   req.Address,
	})

	resp := CreateResponse{
		Success: true,
		Error:   nil,
	}

	if out.Error != nil {
		resp = CreateResponse{
			Success: false,
			Error:   out.Error,
		}
	}

	u.OutputJSON(w, resp)
}

func (u *User) GetByID(w http.ResponseWriter, r *http.Request) {
	// var req GetByIDRequest
	// if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
	// 	log.Println(LogErrDecodeReq)
	// 	u.ErrorBadRequest(w, err)
	// }

	paramID := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(paramID, 0, 64)
	if err != nil {
		u.ErrorBadRequest(w, err)
	}

	out := u.service.GetByID(r.Context(), service.GetByIDIn{
		UserID: int(id),
	})

	if out.Error != nil {
		log.Println("Ошибка:", out.Error.Error())
		u.OutputJSON(w, GetByIDResponse{Success: false, Error: out.Error})
		return
	}

	u.OutputJSON(w, GetByIDResponse{Success: true, User: out.User})
}

func (u *User) Update(w http.ResponseWriter, r *http.Request) {
	var req UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(LogErrDecodeReq)
		u.ErrorBadRequest(w, err)
	}

	out := u.service.Update(r.Context(), service.UpdateIn{
		ID:        req.ID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Username:  req.Username,
		Email:     req.Email,
		Address:   req.Address,
	})

	if out.Error != nil {
		u.OutputJSON(w, UpdateResponse{Success: false, Error: out.Error})
		return
	}

	u.OutputJSON(w, UpdateResponse{Success: true, Error: nil})
}

func (u *User) Delete(w http.ResponseWriter, r *http.Request) {
	var req DeleteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(LogErrDecodeReq)
		u.ErrorBadRequest(w, err)
	}

	out := u.service.Delete(r.Context(), service.DeleteIn{
		TableName: req.TableName,
		UserID:    req.ID,
	})

	if out.Error != nil {
		u.OutputJSON(w, DeleteResponse{Success: false, Error: out.Error})
		return
	}

	u.OutputJSON(w, DeleteResponse{Success: true, Error: nil})
}

func (u *User) List(w http.ResponseWriter, r *http.Request) {
	var req ListRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(LogErrDecodeReq)
		u.ErrorBadRequest(w, err)
	}

	out := u.service.List(r.Context(), service.ListIn{
		Limit:  req.Limit,
		Offset: req.Offset,
	})

	// убрать лог
	if out.Error != nil {
		log.Println("ОШИБКА:", out.Error.Error())
		u.OutputJSON(w, ListResponse{Success: false, Error: out.Error})
		return
	}

	u.OutputJSON(w, ListResponse{Success: true, Users: out.Users})
}
