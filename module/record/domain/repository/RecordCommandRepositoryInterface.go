package repository

import (
	"gomora/module/record/domain/entity"
	"gomora/module/record/infrastructure/repository/types"
)

// RecordCommandRepositoryInterface holds the implementable methods for record command repository
type RecordCommandRepositoryInterface interface {
	InsertRecord(data types.CreateRecord) (entity.Record, error)
}
