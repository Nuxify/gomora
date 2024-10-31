package repository

import (
	"gomora/module/record/domain/entity"
)

// RecordQueryRepositoryInterface holds the implementable method for record query repository
type RecordQueryRepositoryInterface interface {
	// SelectRecordByID gets a record by its ID
	SelectRecordByID(ID string) (entity.Record, error)
}
