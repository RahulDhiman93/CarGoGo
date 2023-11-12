package handlers

import (
	"InrixBackend/internal/driver"
	"InrixBackend/internal/models"
	_ "InrixBackend/internal/models"
	"InrixBackend/internal/repository"
	"InrixBackend/internal/repository/dbrepo"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	_ "github.com/go-chi/chi"
	"net/http"
)

var Repo *Repository

type Repository struct {
	DB repository.DatabaseRepo
}

// NewRepo creates a new repo
func NewRepo(db *driver.DB) *Repository {
	return &Repository{
		DB: dbrepo.NewPostgresRepo(db.SQL),
	}
}

// FreshHandlers sets the repository for the handlers
func FreshHandlers(r *Repository) {
	Repo = r
}

type jsonResponse struct {
	OK      bool                   `json:"ok"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

// RegisterUser creates a new user
func (m *Repository) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var requestBody models.AuthRequestBody
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		internalServerError(w, err)
		return
	}

	email := requestBody.Email
	password := requestBody.Password
	firstName := requestBody.FirstName
	lastName := requestBody.LastName
	phone := requestBody.Phone

	if email == "" || password == "" || firstName == "" || lastName == "" {
		internalServerError(w, fmt.Errorf("please add all attributes"))
		return
	}

	token, err := m.DB.RegisterUser(email, password, firstName, lastName, phone)
	if err != nil {
		internalServerError(w, err)
		return
	}
	respData := make(map[string]interface{})
	respData["token"] = token

	response := jsonResponse{
		OK:      true,
		Message: "user registered successfully",
		Data:    respData,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// LoginUser creates a new user
func (m *Repository) LoginUser(w http.ResponseWriter, r *http.Request) {
	var requestBody models.AuthRequestBody
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		internalServerError(w, err)
		return
	}

	email := requestBody.Email
	password := requestBody.Password

	if email == "" || password == "" {
		internalServerError(w, fmt.Errorf("please add all attributes"))
		return
	}

	user, err := m.DB.LoginUser(email, password)
	if err != nil {
		internalServerError(w, err)
		return
	}

	respData := make(map[string]interface{})
	respData["user"] = user
	response := jsonResponse{
		OK:      true,
		Message: "user registered successfully",
		Data:    respData,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// AccessTokenLogin creates a new user
func (m *Repository) AccessTokenLogin(w http.ResponseWriter, r *http.Request) {
	token := chi.URLParam(r, "access_token")
	if token == "" {
		internalServerError(w, fmt.Errorf("please add access_token"))
		return
	}

	user, err := m.DB.AccessTokenLogin(token)
	if err != nil {
		internalServerError(w, err)
		return
	}

	respData := make(map[string]interface{})
	respData["user"] = user
	response := jsonResponse{
		OK:      true,
		Message: "user registered successfully",
		Data:    respData,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
