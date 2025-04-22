package types

import "time"

// CreateRecord service types for create record
type CreateRecord struct {
	ID   string
	Data string
}

type JWTResponse struct {
	AccessToken string
	ExpiresAt   time.Time
}
