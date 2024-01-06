package handlers

import (
	"CarGoGo/internal/driver"
	"CarGoGo/internal/models"
	_ "CarGoGo/internal/models"
	"CarGoGo/internal/repository"
	"CarGoGo/internal/repository/dbrepo"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	_ "github.com/go-chi/chi"
	"net/http"
	"strconv"
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

	user, err := m.DB.RegisterUser(email, password, firstName, lastName, phone)
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
		Message: "user authenticated successfully",
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
		internalServerError(w, fmt.Errorf("user not found"))
		return
	}

	respData := make(map[string]interface{})
	respData["user"] = user
	response := jsonResponse{
		OK:      true,
		Message: "user authenticated successfully",
		Data:    respData,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// PostRide Creates a new ride on platform
func (m *Repository) PostRide(w http.ResponseWriter, r *http.Request) {
	var requestBody models.Ride
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		internalServerError(w, err)
		return
	}

	userId := requestBody.UserId
	license := requestBody.License
	dl := requestBody.DL
	dateTime := requestBody.DateTime
	price := requestBody.Price
	fromAddress := requestBody.FromAddress
	toAddress := requestBody.ToAddress

	if userId == 0 || fromAddress == "" || toAddress == "" || price == 0 || dateTime == "" || license == "" || dl == "" {
		internalServerError(w, fmt.Errorf("please add all attributes"))
		return
	}

	err = m.DB.PostRide(requestBody)
	if err != nil {
		internalServerError(w, err)
		return
	}

	response := jsonResponse{
		OK:      true,
		Message: "ride posted successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetRideInfo gets the info of the ride by ID
func (m *Repository) GetRideInfo(w http.ResponseWriter, r *http.Request) {
	ride_id := chi.URLParam(r, "id")
	if ride_id == "" {
		internalServerError(w, fmt.Errorf("please add ride_id"))
		return
	}

	rideId, err := strconv.Atoi(ride_id)
	if err != nil {
		internalServerError(w, fmt.Errorf("ride not found"))
		return
	}
	ride, err := m.DB.GetRideInfo(rideId)
	if err != nil {
		internalServerError(w, fmt.Errorf("ride not found"))
		return
	}

	respData := make(map[string]interface{})
	respData["ride"] = ride
	response := jsonResponse{
		OK:      true,
		Message: "ride fetched successfully",
		Data:    respData,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// RaiseRideRequest Creates a new ride on platform
func (m *Repository) RaiseRideRequest(w http.ResponseWriter, r *http.Request) {
	var requestBody models.RaiseRideRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		internalServerError(w, err)
		return
	}

	rideId := requestBody.RideId
	reqCustId := requestBody.ReqCustId

	if rideId == 0 || reqCustId == 0 {
		internalServerError(w, fmt.Errorf("please add all attributes"))
		return
	}

	err = m.DB.RaiseRideRequest(requestBody)
	if err != nil {
		internalServerError(w, err)
		return
	}

	response := jsonResponse{
		OK:      true,
		Message: "request raised successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// ConfirmRide confirm the ride from host to customer
func (m *Repository) ConfirmRide(w http.ResponseWriter, r *http.Request) {
	var requestBody models.ConfirmRide
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		internalServerError(w, err)
		return
	}

	rideId := requestBody.RideId
	customerId := requestBody.CustomerId

	if rideId == 0 || customerId == 0 {
		internalServerError(w, fmt.Errorf("please add all attributes"))
		return
	}

	err = m.DB.ConfirmRide(requestBody)
	if err != nil {
		internalServerError(w, err)
		return
	}

	response := jsonResponse{
		OK:      true,
		Message: "ride confirmed successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetRides Gets the rides to customer according to date, from and to
func (m *Repository) GetRides(w http.ResponseWriter, r *http.Request) {
	var requestBody models.GetRides
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		internalServerError(w, err)
		return
	}

	if requestBody.DateTime == "" || requestBody.FromLat == 0.0 || requestBody.FromLong == 0.0 || requestBody.ToLat == 0.0 || requestBody.ToLong == 0.0 {
		internalServerError(w, fmt.Errorf("please add all attributes"))
		return
	}

	rides, err := m.DB.GetRides(requestBody)
	if err != nil {
		internalServerError(w, err)
		return
	}

	respData := make(map[string]interface{})
	respData["rides"] = rides
	response := jsonResponse{
		OK:      true,
		Message: "rides fetched successfully",
		Data:    respData,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
