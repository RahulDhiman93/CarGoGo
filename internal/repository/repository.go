package repository

import "InrixBackend/internal/models"

type DatabaseRepo interface {
	RegisterUser(email, password, firstName, lastName, phone string) (string, error)
	LoginUser(email, password string) (models.User, error)
	AccessTokenLogin(token string) (models.User, error)
}
