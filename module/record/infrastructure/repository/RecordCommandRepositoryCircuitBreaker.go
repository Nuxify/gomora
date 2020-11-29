package repository

import (
	"github.com/afex/hystrix-go/hystrix"

	hystrix_config "gomora/configs/hystrix"
	"gomora/module/record/domain/entity"
	"gomora/module/record/domain/repository"
	repositoryTypes "gomora/module/record/infrastructure/repository/types"
)

// RecordCommandRepositoryCircuitBreaker circuit breaker for record command repository
type RecordCommandRepositoryCircuitBreaker struct {
	repository.RecordCommandRepositoryInterface
}

var config = hystrix_config.Config{}

// InsertRecord decorator pattern to insert record
func (repository *RecordCommandRepositoryCircuitBreaker) InsertRecord(data repositoryTypes.CreateRecord) (entity.Record, error) {
	output := make(chan entity.Record, 1)
	hystrix.ConfigureCommand("insert_record", config.Settings())
	errors := hystrix.Go("insert_record", func() error {
		record, err := repository.RecordCommandRepositoryInterface.InsertRecord(data)
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
