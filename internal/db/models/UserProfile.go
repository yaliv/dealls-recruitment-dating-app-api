package models

import (
	"time"
)

type UserProfile struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    int
	Verified  bool
	Name      *string
	Age       *int
	Bio       *string
	PicUrl    *string
}
