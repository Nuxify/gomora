package types

import "time"

type CreateRecord struct {
	ID   string
	Data string
}

type JWTResponse struct {
	AccessToken string
	ExpiresAt   time.Time
}
