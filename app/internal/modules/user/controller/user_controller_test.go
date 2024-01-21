package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"

	"projects/LDmitryLD/repository/app/internal/infrastructure/responder"
	"projects/LDmitryLD/repository/app/internal/models"
	"projects/LDmitryLD/repository/app/internal/modules/user/service"
	"projects/LDmitryLD/repository/app/internal/modules/user/service/mocks"

	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	testReqErr        = "ошибка при выполнении тестового запроса:"
	testRespDecodeErr = "ошибка при декодировании тестового ответа:"
)

func TestUser_Create_BadRequest(t *testing.T) {
	user := User{
		Responder: &responder.Respond{},
	}

	req := map[string]interface{}{"first_name": 1}
	reqJSON, _ := json.Marshal(req)

	s := httptest.NewServer(http.HandlerFunc(user.Craete))
	defer s.Close()

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(testReqErr, err.Error())
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestUser_Create(t *testing.T) {
	serviceMock := mocks.NewUserer(t)

	req := CreateRequest{
		FirstName: "fitsName",
		LastName:  "lastName",
		Username:  "username",
		Email:     "email",
		Address:   "address",
	}
	in := service.CreateIn{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Username:  req.Username,
		Email:     req.Email,
		Address:   req.Address,
	}
	out := service.CreateOut{
		Error: nil,
	}
	expect := CreateResponse{
		Success: true,
		Error:   "",
	}

	serviceMock.On("Create", mock.Anything, in).Return(out)

	user := NewUser(serviceMock)

	reqJSON, _ := json.Marshal(req)

	s := httptest.NewServer(http.HandlerFunc(user.Craete))
	defer s.Close()

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(testReqErr, err.Error())
	}
	defer resp.Body.Close()

	var createResp CreateResponse
	if err = json.NewDecoder(resp.Body).Decode(&createResp); err != nil {
		t.Fatal(testRespDecodeErr, err.Error())
	}

	assert.Equal(t, expect, createResp)
}

func TestUser_Create_Err(t *testing.T) {
	serviceMock := mocks.NewUserer(t)
	req := CreateRequest{}
	in := service.CreateIn{}
	out := service.CreateOut{
		Error: errors.New("test error"),
	}
	expect := CreateResponse{
		Success: false,
		Error:   "test error",
	}
	serviceMock.On("Create", mock.Anything, in).Return(out)

	user := NewUser(serviceMock)

	reqJSON, _ := json.Marshal(req)

	s := httptest.NewServer(http.HandlerFunc(user.Craete))
	defer s.Close()

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(testReqErr, err.Error())
	}
	defer resp.Body.Close()

	var createResp CreateResponse
	if err := json.NewDecoder(resp.Body).Decode(&createResp); err != nil {
		t.Fatal(testRespDecodeErr, err.Error())
	}

	assert.Equal(t, expect, createResp)
}

func TestUser_GetByID(t *testing.T) {

	in := service.GetByIDIn{
		UserID: 1,
	}
	testUser := models.User{
		ID: 1,
	}
	out := service.GetByIDOut{
		User:  testUser,
		Error: nil,
	}
	expect := GetByIDResponse{
		Success: true,
		User:    testUser,
		Error:   "",
	}
	serviceMock := mocks.NewUserer(t)
	serviceMock.On("GetByID", mock.Anything, in).Return(out)

	user := NewUser(serviceMock)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/users/1", nil)

	r := chi.NewRouter()
	r.MethodFunc("GET", "/api/users/{id}", user.GetByID)

	r.ServeHTTP(rr, req)

	var resp GetByIDResponse
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatal(testRespDecodeErr, err.Error())
	}

	assert.Equal(t, expect, resp)
}

func TestUser_GetByID_BadRequest(t *testing.T) {
	user := User{
		Responder: &responder.Respond{},
	}

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/users/b", nil)

	r := chi.NewRouter()

	r.MethodFunc("GET", "/api/users/{id}", user.GetByID)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)

}

func TestUser_GetByID_Error(t *testing.T) {
	in := service.GetByIDIn{
		UserID: 1,
	}
	out := service.GetByIDOut{
		Error: errors.New("test error"),
	}
	expect := GetByIDResponse{
		Success: false,
		Error:   "test error",
	}

	serviceMock := mocks.NewUserer(t)
	serviceMock.On("GetByID", mock.Anything, in).Return(out)

	user := NewUser(serviceMock)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/users/1", nil)

	r := chi.NewRouter()

	r.MethodFunc("GET", "/api/users/{id}", user.GetByID)

	r.ServeHTTP(rr, req)

	var resp GetByIDResponse
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatal(testRespDecodeErr, err.Error())
	}

	assert.Equal(t, expect, resp)

}

func TestUser_Update(t *testing.T) {
	in := service.UpdateIn{
		ID:        1,
		FirstName: "name",
	}
	out := service.UpdateOut{
		Error: nil,
	}
	req := UpdateRequest{
		ID:        1,
		FirstName: "name",
	}
	expect := UpdateResponse{
		Success: true,
		Error:   "",
	}

	serviceMock := mocks.NewUserer(t)
	serviceMock.On("Update", mock.Anything, in).Return(out)

	user := NewUser(serviceMock)

	reqJSON, _ := json.Marshal(req)

	s := httptest.NewServer(http.HandlerFunc(user.Update))
	defer s.Close()

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(testReqErr, err.Error())
	}
	defer resp.Body.Close()

	var updateResp UpdateResponse
	if err := json.NewDecoder(resp.Body).Decode(&updateResp); err != nil {
		t.Fatal(testRespDecodeErr, err.Error())
	}

	assert.Equal(t, expect, updateResp)
}

func TestUser_Update_BadRequest(t *testing.T) {
	user := User{
		Responder: &responder.Respond{},
	}

	req := map[string]interface{}{"ID": "1"}

	reqJSON, _ := json.Marshal(req)

	s := httptest.NewServer(http.HandlerFunc(user.Update))
	defer s.Close()

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(testReqErr, err.Error())
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestUser_Update_Error(t *testing.T) {
	in := service.UpdateIn{
		ID: 1,
	}
	out := service.UpdateOut{
		Error: errors.New("test error"),
	}
	req := UpdateRequest{
		ID: 1,
	}
	expect := UpdateResponse{
		Success: false,
		Error:   "test error",
	}

	serviceMock := mocks.NewUserer(t)
	serviceMock.On("Update", mock.Anything, in).Return(out)

	user := NewUser(serviceMock)

	s := httptest.NewServer(http.HandlerFunc(user.Update))
	defer s.Close()

	reqJSON, _ := json.Marshal(req)

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(testReqErr, err.Error())
	}
	defer resp.Body.Close()

	var updateResp UpdateResponse
	if err := json.NewDecoder(resp.Body).Decode(&updateResp); err != nil {
		t.Fatal(testRespDecodeErr, err.Error())
	}

	assert.Equal(t, expect, updateResp)
}

func TestUser_Delete(t *testing.T) {
	in := service.DeleteIn{
		TableName: "users",
		UserID:    1,
	}
	out := service.DeleteOut{
		Error: nil,
	}
	req := DeleteRequest{
		TableName: "users",
		ID:        1,
	}
	expect := DeleteResponse{
		Success: true,
		Error:   "",
	}

	serviceMock := mocks.NewUserer(t)
	serviceMock.On("Delete", mock.Anything, in).Return(out)

	user := NewUser(serviceMock)

	s := httptest.NewServer(http.HandlerFunc(user.Delete))
	defer s.Close()

	reqJSON, _ := json.Marshal(req)

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(testReqErr, err.Error())
	}
	defer resp.Body.Close()

	var deleteResp DeleteResponse
	if err := json.NewDecoder(resp.Body).Decode(&deleteResp); err != nil {
		t.Fatal(testRespDecodeErr)
	}

	assert.Equal(t, expect, deleteResp)
}

func TestUser_Delete_BadRequest(t *testing.T) {
	user := User{
		Responder: &responder.Respond{},
	}

	req := map[string]interface{}{"id": "1"}
	reqJSON, _ := json.Marshal(req)

	s := httptest.NewServer(http.HandlerFunc(user.Delete))
	defer s.Close()

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(testReqErr, err.Error())
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestUser_Delete_Error(t *testing.T) {
	in := service.DeleteIn{
		TableName: "pets",
		UserID:    1,
	}
	out := service.DeleteOut{
		Error: errors.New("test error"),
	}
	req := DeleteRequest{
		TableName: "pets",
		ID:        1,
	}
	expect := DeleteResponse{
		Success: false,
		Error:   "test error",
	}

	serviceMock := mocks.NewUserer(t)
	serviceMock.On("Delete", mock.Anything, in).Return(out)

	user := NewUser(serviceMock)

	s := httptest.NewServer(http.HandlerFunc(user.Delete))
	defer s.Close()

	reqJSON, _ := json.Marshal(req)

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(testReqErr)
	}
	defer resp.Body.Close()

	var deleteResp DeleteResponse
	if err := json.NewDecoder(resp.Body).Decode(&deleteResp); err != nil {
		t.Fatal(testRespDecodeErr)
	}

	assert.Equal(t, expect, deleteResp)
}

func TestUser_List(t *testing.T) {
	users := []models.User{{ID: 1}, {ID: 2}}
	in := service.ListIn{
		Limit:  3,
		Offset: 1,
	}
	out := service.ListOut{
		Users: users,
		Error: nil,
	}
	req := ListRequest{
		Limit:  3,
		Offset: 1,
	}
	expect := ListResponse{
		Success: true,
		Error:   "",
		Users:   users,
	}

	serviceMock := mocks.NewUserer(t)
	serviceMock.On("List", mock.Anything, in).Return(out)

	user := NewUser(serviceMock)

	s := httptest.NewServer(http.HandlerFunc(user.List))
	defer s.Close()

	reqJSON, _ := json.Marshal(req)

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(testReqErr, err.Error())
	}
	defer resp.Body.Close()

	var listResp ListResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		t.Fatal(testRespDecodeErr, err.Error())
	}

	assert.Equal(t, expect, listResp)
}

func TestUser_List_BadRequest(t *testing.T) {
	user := User{
		Responder: &responder.Respond{},
	}

	req := map[string]interface{}{"limit": "1"}
	reqJSON, _ := json.Marshal(req)

	s := httptest.NewServer(http.HandlerFunc(user.List))
	defer s.Close()

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(testReqErr, err.Error())
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestUser_List_Error(t *testing.T) {
	in := service.ListIn{
		Limit:  3,
		Offset: 2,
	}
	out := service.ListOut{
		Error: errors.New("test error"),
	}
	req := ListRequest{
		Limit:  3,
		Offset: 2,
	}
	expect := ListResponse{
		Success: false,
		Error:   "test error",
	}

	serviceMock := mocks.NewUserer(t)
	serviceMock.On("List", mock.Anything, in).Return(out)

	user := NewUser(serviceMock)

	s := httptest.NewServer(http.HandlerFunc(user.List))
	defer s.Close()

	reqJSON, _ := json.Marshal(req)

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Error(testReqErr, err.Error())
	}
	defer resp.Body.Close()

	var listResp ListResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		t.Fatal(testRespDecodeErr, err.Error())
	}

	assert.Equal(t, expect, listResp)
}
