package domain

import "time"

type Role string

const (
	Reader Role = "reader"
	Agent  Role = "agent"
	Admin  Role = "admin"
)

type User struct {
	ID        uint64
	Name      string
	Email     string
	Password  string
	Role      Role
	CreatedAt time.Time
	UpdatedAt time.Time
}
