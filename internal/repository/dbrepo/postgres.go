package dbrepo

import (
	"InrixBackend/internal/models"
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

// RegisterUser Register the user
func (m *postgresDBRepo) RegisterUser(email, password, firstName, lastName, phone string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var existingEmail string
	err := m.DB.QueryRowContext(ctx, "SELECT email FROM users WHERE email = $1", email).Scan(&existingEmail)
	if err == nil {
		return "", fmt.Errorf("email %s is already registered", email)
	} else if !errors.Is(err, sql.ErrNoRows) {
		log.Println("Error checking existing email:", err)
		return "", err
	}

	query := `insert into users (access_token, first_name, last_name, email, password, phone,created_at,updated_at) values ($1, $2, $3, $4, $5, $6, $7, $8)`

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error generating bcrypt hash:", err)
		return "", err
	}

	randomBytes := make([]byte, 32)
	_, err = rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	token := base64.URLEncoding.EncodeToString(randomBytes)

	_, err = m.DB.ExecContext(ctx, query, token, firstName, lastName, email, hashedPassword, phone, time.Now(), time.Now())

	if err != nil {
		log.Println(err)
		return "", err
	}
	return token, nil
}

// LoginUser Login the user
func (m *postgresDBRepo) LoginUser(email, password string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var user models.User

	query := `SELECT id, access_token, first_name, last_name, email, password, phone, access_level FROM users WHERE email = $1`

	err := m.DB.QueryRowContext(ctx, query, email).Scan(
		&user.Id,
		&user.AccessToken,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.Phone,
		&user.AccessLevel,
	)
	if err != nil {
		return user, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return user, errors.New("incorrect password")
	}
	return user, nil
}

// AccessTokenLogin Login the user with token
func (m *postgresDBRepo) AccessTokenLogin(token string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var user models.User

	query := `SELECT id, access_token, first_name, last_name, email, password, phone, access_level FROM users WHERE access_token = $1`

	err := m.DB.QueryRowContext(ctx, query, token).Scan(
		&user.Id,
		&user.AccessToken,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.Phone,
		&user.AccessLevel,
	)
	if err != nil {
		return user, err
	}

	return user, nil
}
