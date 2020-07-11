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
	tenantREST "gomora/module/tenant/interfaces/http/rest"
)

// ServiceContainerInterface contains the dependency injected instances
type ServiceContainerInterface interface {
	// gRPC
	RegisterTenantGRPCQueryController() tenantGRPC.TenantQueryController

	// REST
	RegisterTenantRESTQueryController() tenantREST.TenantQueryController
}

type kernel struct{}

var (
	m              sync.Mutex
	k              *kernel
	containerOnce  sync.Once
	mysqlDBHandler *mysql.MySQLDBHandler
)

//================================= gRPC ===================================
// RegisterTenantGRPCQueryController performs dependency injection to the RegisterTenantGRPCQueryController
func (k *kernel) RegisterTenantGRPCQueryController() tenantGRPC.TenantQueryController {
	service := k.tenantQueryServiceContainer()

	controller := tenantGRPC.TenantQueryController{
		TenantQueryServiceInterface: service,
	}

	return controller
}

//==========================================================================

//================================= REST ===================================
// RegisterTenantRESTQueryController performs dependency injection to the RegisterTenantRESTQueryController
func (k *kernel) RegisterTenantRESTQueryController() tenantREST.TenantQueryController {
	service := k.tenantQueryServiceContainer()

	controller := tenantREST.TenantQueryController{
		TenantQueryServiceInterface: service,
	}

	return controller
}

//==========================================================================

func (k *kernel) tenantQueryServiceContainer() *tenantService.TenantQueryService {
	repository := &tenantRepository.TenantQueryRepository{
		MySQLDBHandlerInterface: mysqlDBHandler,
	}

	service := &tenantService.TenantQueryService{
		TenantQueryRepositoryInterface: &tenantRepository.TenantQueryRepositoryCircuitBreaker{
			TenantQueryRepositoryInterface: repository,
		},
	}

	return service
}

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
	m.Lock()
	defer m.Unlock()

	if k == nil {
		containerOnce.Do(func() {
			// register container handlers
			registerHandlers()

			k = &kernel{}
		})
	}
	return k
}
