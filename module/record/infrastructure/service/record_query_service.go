package service

import (
	"context"

	"gomora/module/record/domain/entity"
	"gomora/module/record/domain/repository"
)

// RecordQueryService handles the record query service logic
type RecordQueryService struct {
	repository.RecordQueryRepositoryInterface
}

// GetRecordByID retrieves the record provided by its id
func (service *RecordQueryService) GetRecordByID(ctx context.Context, ID string) (entity.Record, error) {
	res, err := service.RecordQueryRepositoryInterface.SelectRecordByID(ID)
	if err != nil {
		return res, err
	}

	return res, nil
}
