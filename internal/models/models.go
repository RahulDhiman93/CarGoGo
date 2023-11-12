package models

type AuthRequestBody struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
}

type Ride struct {
	User    User
	CarType int
	License string
	DL      string
	From    location
	To      location
	Price   int
	Date    string
	Time    string
}

type location struct {
	Address   string
	Latitude  float64
	Longitude float64
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
