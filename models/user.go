package models

import "time"

type User struct {
	Id          int       `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Role        string    `json:"role"`
	About       string    `json:"about"`
	Avatar      string    `json:"avatar"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	Gender      string    `json:"gender"`
	Postcode    int       `json:"postcode"`
	Birthday    time.Time `json:"birthday"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
