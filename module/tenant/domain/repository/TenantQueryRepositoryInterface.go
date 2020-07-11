package repository

import (
	"gomora/module/tenant/domain/entity"
)

type TenantQueryRepositoryInterface interface {
	SelectTenantByID(tenantID string) (entity.Tenant, error)
}
