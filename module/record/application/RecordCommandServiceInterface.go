package application

import (
	"context"

	"gomora/module/record/domain/entity"
	"gomora/module/record/infrastructure/service/types"
)

// RecordCommandServiceInterface holds the implementable methods for the record command service
type RecordCommandServiceInterface interface {
	// CreateRecord creates a new record
	CreateRecord(ctx context.Context, data types.CreateRecord) (entity.Record, error)
	// GenerateToken generates a jwt token
	GenerateToken(ctx context.Context) (types.JWTResponse, error)
}
