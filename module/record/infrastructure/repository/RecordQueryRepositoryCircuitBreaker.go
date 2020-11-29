package repository

import (
	"github.com/afex/hystrix-go/hystrix"

	"gomora/module/record/domain/entity"
	"gomora/module/record/domain/repository"
)

// RecordQueryRepositoryCircuitBreaker holds the implementable methods for record query circuitbreaker
type RecordQueryRepositoryCircuitBreaker struct {
	repository.RecordQueryRepositoryInterface
}

// SelectRecordByID decorator pattern for select record repository
func (repository *RecordQueryRepositoryCircuitBreaker) SelectRecordByID(ID string) (entity.Record, error) {
	output := make(chan entity.Record, 1)
	hystrix.ConfigureCommand("select_record_by_id", config.Settings())
	errors := hystrix.Go("select_record_by_id", func() error {
		record, err := repository.RecordQueryRepositoryInterface.SelectRecordByID(ID)
		if err != nil {
			return err
		}

		output <- record
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		return entity.Record{}, err
	}
}
