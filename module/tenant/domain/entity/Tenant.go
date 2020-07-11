package entity

import (
	"time"
)

// Tenant holds the tenant entity fields
type Tenant struct {
	ID            string
	Name          string
	Alias         string
	Email         string
	Code          string
	Address       string
	ContactNumber string `db:"contact_number"`
	Avatar        string
	IsActive      bool      `db:"is_active"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}

// GetModelName returns the model name of tenant entity that can be used for naming schemas
func (entity *Tenant) GetModelName() string {
	return "tenants"
}
