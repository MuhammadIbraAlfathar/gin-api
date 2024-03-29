package entity

import "time"

type User struct {
	ID        int
	Name      string
	Email     string
	Password  string
	Gender    string
	UpdatedAt time.Time
	CreatedAt time.Time
}
