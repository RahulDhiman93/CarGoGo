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

// PostRide Post the ride by host to users
func (m *postgresDBRepo) PostRide(r models.Ride) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	timeString := r.DateTime
	layout := "2006-01-02 15:04"
	timestamp, err := time.Parse(layout, timeString)
	if err != nil {
		return err
	}
	query := `insert into rides (user_id, car_type, license_plate, dl, date_time, price, status, from_address, from_lat, from_long, to_address, to_lat, to_long,created_at,updated_at) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)`

	_, err = m.DB.ExecContext(ctx, query,
		r.UserId,
		r.CarType,
		r.License,
		r.DL,
		timestamp,
		r.Price,
		r.Status,
		r.FromAddress,
		r.FromLat,
		r.FromLong,
		r.ToAddress,
		r.ToLat,
		r.ToLong,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// GetRideInfo gets the ride info from DB
func (m *postgresDBRepo) GetRideInfo(id int) (models.Ride, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var ride models.Ride

	query := `SELECT id, user_id, car_type, license_plate, dl, date_time, price, status, req_cust_ids, customer_id, from_address, from_lat, from_long, to_address, to_lat, to_long FROM rides WHERE id = $1`

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&ride.ID,
		&ride.UserId,
		&ride.CarType,
		&ride.License,
		&ride.DL,
		&ride.DateTime,
		&ride.Price,
		&ride.Status,
		&ride.ReqCustIds,
		&ride.CustomerId,
		&ride.FromAddress,
		&ride.FromLat,
		&ride.FromLong,
		&ride.ToAddress,
		&ride.ToLat,
		&ride.ToLong,
	)
	if err != nil {
		return ride, err
	}

	return ride, nil
}

// RaiseRideRequest Raise a request to ride by users to host
func (m *postgresDBRepo) RaiseRideRequest(r models.RaiseRideRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `UPDATE rides SET req_cust_ids = ARRAY_APPEND(req_cust_ids, $2) WHERE id = $1;`

	_, err := m.DB.ExecContext(ctx, query,
		r.RideId,
		r.ReqCustId,
	)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
