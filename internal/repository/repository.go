package repository

import "CarGoGo/internal/models"

type DatabaseRepo interface {
	RegisterUser(email, password, firstName, lastName, phone string) (string, error)
	LoginUser(email, password string) (models.User, error)
	AccessTokenLogin(token string) (models.User, error)
	PostRide(r models.Ride) error
	GetRideInfo(id int) (models.Ride, error)
	RaiseRideRequest(r models.RaiseRideRequest) error
	ConfirmRide(r models.ConfirmRide) error
	GetRides(r models.GetRides) ([]models.Ride, error)
}
