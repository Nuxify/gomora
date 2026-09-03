package repository

import (
	"database/sql"
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

	stmt := fmt.Sprintf("SELECT * FROM %s WHERE id=:id", record.GetModelName())
	err := repository.QueryRow(stmt, map[string]interface{}{
		"id": ID,
	}, &record)
	if err != nil {
		if err == sql.ErrNoRows {
			return record, errors.New(apiError.MissingRecord)
		}

		return record, errors.New(apiError.DatabaseError)
	}

	return record, nil
}
