package handlers

import (
	"InrixBackend/internal/driver"
	"InrixBackend/internal/forms"
	_ "InrixBackend/internal/models"
	"InrixBackend/internal/repository"
	"InrixBackend/internal/repository/dbrepo"
	"encoding/json"
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
	OK       bool   `json:"ok"`
	Message  string `json:"message"`
	LongURL  string `json:"long_url"`
	ShortURL string `json:"short_url"`
	QRCode   []byte `json:"qrcode"`
}

// ShortenURL generates a short key for a given URL and stores it in the map.
func (m *Repository) ShortenURL(w http.ResponseWriter, r *http.Request) {

	response := jsonResponse{
		OK:      true,
		Message: "Short URL created",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// RegisterUser creates a new user
func (m *Repository) RegisterUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		internalServerError(w)
		return
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")
	firstName := r.Form.Get("first_name")
	lastName := r.Form.Get("last_name")
	phone := r.Form.Get("phone")

	form := forms.NewForm(r.PostForm)
	form.Required("email", "password", "name")
	form.IsEmail("email")

	if !form.Valid() {
		internalServerError(w)
		return
	}

	id, _, err := m.DB.Authenticate(email, password)
}
