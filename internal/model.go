package internal

import "time"

type BaseModel struct {
	ID        int64
	CreatedAt *time.Time
	UpdatedAt *time.Time
	// DeletedAt *time.Time
}
