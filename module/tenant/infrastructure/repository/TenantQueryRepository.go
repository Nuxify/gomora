package repository

import (
	"errors"
	"fmt"

	"gomora/infrastructures/database/mysql/types"
	apiError "gomora/internal/errors"
	"gomora/module/tenant/domain/entity"
)

// TenantQueryRepository handles the tenant query repository logic
type TenantQueryRepository struct {
	types.MySQLDBHandlerInterface
}

// SelectTenantByID selects the tenant record by tenant id
func (repository *TenantQueryRepository) SelectTenantByID(tenantID string) (entity.Tenant, error) {
	var tenant entity.Tenant
	var tenants []entity.Tenant

	stmt := fmt.Sprintf("SELECT * FROM %s WHERE id = :id", tenant.GetModelName())
	err := repository.Query(stmt, map[string]interface{}{
		"id": tenantID,
	}, &tenants)
	if err != nil {
		return tenant, errors.New(apiError.DatabaseError)
	} else if len(tenants) == 0 {
		return tenant, errors.New(apiError.MissingRecord)
	}

	return tenants[0], nil
}
