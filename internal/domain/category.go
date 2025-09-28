package domain

import (
	"time"
)

type Category struct {
	ID          uint
	Name        string
	Description string
	Unit        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
