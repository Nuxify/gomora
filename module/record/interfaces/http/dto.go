package http

import (
	"github.com/go-playground/validator/v10"
)

var (
	Validate         *validator.Validate = validator.New(validator.WithRequiredStructEnabled())
	ValidationErrors map[string]string   = map[string]string{
		"CreateRecordRequest.ID":   "ID field is required.",
		"CreateRecordRequest.Data": "Data field is required.",
	}
)

type CreateRecordRequest struct {
	ID   string `json:"id" validate:"required"`
	Data string `json:"data" validate:"required"`
}

type CreateRecordResponse struct {
	ID        string `json:"id"`
	Data      string `json:"data"`
	CreatedAt int64  `json:"createdAt"`
}

type GetRecordResponse struct {
	ID        string `json:"id"`
	Data      string `json:"data"`
	CreatedAt int64  `json:"createdAt"`
}

type GenerateTokenResponse struct {
	AccessToken string `json:"accessToken"`
}
