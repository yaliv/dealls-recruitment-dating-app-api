package models

import (
	"time"
)

type Users struct {
	ID            int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeactivatedAt *time.Time
	Email         string
	Secret        string
}
