package entity

import "time"

type User struct {
	UserId       int       `json:"user_id"`
	UserFullname string    `json:"user_fullname"`
	UserUsername string    `json:"user_username"`
	UserPassword string    `json:"user_password"`
	UserEmail    string    `json:"user_email"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}
