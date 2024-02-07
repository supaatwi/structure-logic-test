package model

import (
	"time"
)

type Todo struct {
	ID          int
	Description string
	CreatedAt   time.Time
}
