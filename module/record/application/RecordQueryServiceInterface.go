package application

import (
	"context"

	"gomora/module/record/domain/entity"
)

// RecordQueryServiceInterface holds the implementable methods for the record query service
type RecordQueryServiceInterface interface {
	// GetRecordByID gets a record by its ID
	GetRecordByID(ctx context.Context, ID string) (entity.Record, error)
}
