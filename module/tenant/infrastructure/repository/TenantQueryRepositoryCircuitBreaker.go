package repository

import (
	"github.com/afex/hystrix-go/hystrix"

	hystrix_config "gomora/configs/hystrix"
	"gomora/module/tenant/domain/entity"
	"gomora/module/tenant/domain/repository"
)

type TenantQueryRepositoryCircuitBreaker struct {
	repository.TenantQueryRepositoryInterface
}

var config = hystrix_config.Config{}

// SelectTenantByID retrieve tenant data by tenant id
func (repository *TenantQueryRepositoryCircuitBreaker) SelectTenantByID(tenantID string) (entity.Tenant, error) {
	output := make(chan entity.Tenant, 1)
	hystrix.ConfigureCommand("select_tenant_by_id", config.Settings())
	errors := hystrix.Go("select_tenant_by_id", func() error {
		tenantModel, err := repository.TenantQueryRepositoryInterface.SelectTenantByID(tenantID)
		if err != nil {
			return err
		}

		output <- tenantModel
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		return entity.Tenant{}, err
	}
}
