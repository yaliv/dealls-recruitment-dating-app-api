package models

import (
	"time"
)

type User struct {
	ID            int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeactivatedAt *time.Time
	Email         string
	Secret        string
}
