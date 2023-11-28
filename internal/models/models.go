package models

type AuthRequestBody struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
}

type Ride struct {
	ID          int         `json:"id"`
	UserId      int         `json:"user_id"`
	CarType     int         `json:"car_type"`
	License     string      `json:"license_plate"`
	DL          string      `json:"dl"`
	FromAddress string      `json:"from_address"`
	FromLat     float64     `json:"from_lat"`
	FromLong    float64     `json:"from_long"`
	ToAddress   string      `json:"to_address"`
	ToLat       float64     `json:"to_lat"`
	ToLong      float64     `json:"to_long"`
	Price       int         `json:"price"`
	Status      int         `json:"status"`
	ReqCustIds  interface{} `json:"req_cust_ids"`
	CustomerId  int         `json:"customer_id"`
	DateTime    string      `json:"date_time"`
}

type RaiseRideRequest struct {
	RideId    int `json:"ride_id"`
	ReqCustId int `json:"req_cust_id"`
}

type ConfirmRide struct {
	RideId     int `json:"ride_id"`
	CustomerId int `json:"customer_id"`
}

type GetRides struct {
	DateTime string  `json:"date_time"`
	FromLat  float64 `json:"from_lat"`
	FromLong float64 `json:"from_long"`
	ToLat    float64 `json:"to_lat"`
	ToLong   float64 `json:"to_long"`
}

type User struct {
	Id          int    `json:"id"`
	AccessToken string `json:"access_token"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Phone       string `json:"phone"`
	AccessLevel int    `json:"access_level"`
}
