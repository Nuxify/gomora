/*
|--------------------------------------------------------------------------
| Service Container
|--------------------------------------------------------------------------
|
| This file performs the compiled dependency injection for your middlewares,
| controllers, services, providers, repositories, etc..
|
*/
package interfaces

import (
	"log"
	"os"
	"sync"

	"gomora/infrastructures/database/mysql"

	tenantRepository "gomora/module/tenant/infrastructure/repository"
	tenantService "gomora/module/tenant/infrastructure/service"
	tenantGRPC "gomora/module/tenant/interfaces/http/grpc"
)

// ServiceContainerInterface contains the dependency injected instances
type ServiceContainerInterface interface {
	RegisterTenantGRPCQueryController() tenantGRPC.TenantQueryController
}

type kernel struct{}

var (
	k              *kernel
	containerOnce  sync.Once
	mysqlDBHandler *mysql.MySQLDBHandler
)

//================================= gRPC ===================================
// RegisterTenantGRPCQueryController performs dependency injection to the RegisterTenantGRPCQueryController
func (k *kernel) RegisterTenantGRPCQueryController() tenantGRPC.TenantQueryController {
	repository := &tenantRepository.TenantQueryRepository{
		MySQLDBHandlerInterface: mysqlDBHandler,
	}

	service := &tenantService.TenantQueryService{
		TenantQueryRepositoryInterface: &tenantRepository.TenantQueryRepositoryCircuitBreaker{
			TenantQueryRepositoryInterface: repository,
		},
	}

	controller := tenantGRPC.TenantQueryController{
		TenantQueryServiceInterface: service,
	}

	return controller
}

//==========================================================================

func registerHandlers() {
	// create new mysql database connection
	mysqlDBHandler = &mysql.MySQLDBHandler{}
	err := mysqlDBHandler.Connect(os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"))
	if err != nil {
		log.Fatalf("[SERVER] mysql database is not responding %v", err)
	}
}

// ServiceContainer export instantiated service container once
func ServiceContainer() ServiceContainerInterface {
	if k == nil {
		containerOnce.Do(func() {
			// register container handlers
			registerHandlers()

			k = &kernel{}
		})
	}
	return k
}
