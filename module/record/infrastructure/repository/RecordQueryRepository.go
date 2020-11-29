package repository

import (
	"errors"
	"fmt"

	"gomora/infrastructures/database/mysql/types"
	apiError "gomora/internal/errors"
	"gomora/module/record/domain/entity"
)

// RecordQueryRepository handles the record query repository logic
type RecordQueryRepository struct {
	types.MySQLDBHandlerInterface
}

// SelectRecordByID select a record by id
func (repository *RecordQueryRepository) SelectRecordByID(ID string) (entity.Record, error) {
	var record entity.Record
	var records []entity.Record

	stmt := fmt.Sprintf("SELECT * FROM %s WHERE id=:id", record.GetModelName())
	err := repository.Query(stmt, map[string]interface{}{
		"id": ID,
	}, &records)
	if err != nil {
		return record, errors.New(apiError.DatabaseError)
	} else if len(records) == 0 {
		return record, errors.New(apiError.MissingRecord)
	}

	return records[0], nil
}
