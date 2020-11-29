package entity

import (
	"time"
)

// Record holds the record entity fields
type Record struct {
	ID        string
	Data      string
	CreatedAt time.Time `db:"created_at"`
}

// GetModelName returns the model name of record entity that can be used for naming schemas
func (entity *Record) GetModelName() string {
	return "records"
}
