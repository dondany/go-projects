package model

import "time"

type User struct {
	ID        int32
	Email     string
	Name      string
	Password  string
	CreatedAt time.Time
	Token string
}

type UserLogin struct {
	Email    string
	Password string
}