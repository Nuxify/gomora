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
	"gomora/infrastructures/database/mysql/types"
	recordRepository "gomora/module/record/infrastructure/repository"
	recordService "gomora/module/record/infrastructure/service"
	recordGRPC "gomora/module/record/interfaces/http/grpc"
	recordREST "gomora/module/record/interfaces/http/rest"
)

// ServiceContainerInterface contains the dependency injected instances
type ServiceContainerInterface interface {
	// gRPC
	RegisterRecordGRPCCommandController() recordGRPC.RecordCommandController
	RegisterRecordGRPCQueryController() recordGRPC.RecordQueryController

	// REST
	RegisterRecordRESTCommandController() recordREST.RecordCommandController
	RegisterRecordRESTQueryController() recordREST.RecordQueryController
}

type kernel struct{}

var (
	m              sync.Mutex
	k              *kernel
	containerOnce  sync.Once
	mysqlDBHandler *mysql.MySQLDBHandler
)

// ================================= gRPC ===================================
// RegisterRecordGRPCCommandController performs dependency injection to the RegisterRecordGRPCCommandController
func (k *kernel) RegisterRecordGRPCCommandController() recordGRPC.RecordCommandController {
	service := k.recordCommandServiceContainer()

	controller := recordGRPC.RecordCommandController{
		RecordCommandServiceInterface: service,
	}

	return controller
}

// RegisterRecordGRPCQueryController performs dependency injection to the RegisterRecordGRPCQueryController
func (k *kernel) RegisterRecordGRPCQueryController() recordGRPC.RecordQueryController {
	service := k.recordQueryServiceContainer()

	controller := recordGRPC.RecordQueryController{
		RecordQueryServiceInterface: service,
	}

	return controller
}

//==========================================================================

// ================================= REST ===================================
// RegisterRecordRESTCommandController performs dependency injection to the RegisterRecordRESTCommandController
func (k *kernel) RegisterRecordRESTCommandController() recordREST.RecordCommandController {
	service := k.recordCommandServiceContainer()

	controller := recordREST.RecordCommandController{
		RecordCommandServiceInterface: service,
	}

	return controller
}

// RegisterRecordRESTQueryController performs dependency injection to the RegisterRecordRESTQueryController
func (k *kernel) RegisterRecordRESTQueryController() recordREST.RecordQueryController {
	service := k.recordQueryServiceContainer()

	controller := recordREST.RecordQueryController{
		RecordQueryServiceInterface: service,
	}

	return controller
}

//==========================================================================

func (k *kernel) recordCommandServiceContainer() *recordService.RecordCommandService {
	repository := &recordRepository.RecordCommandRepository{
		MySQLDBHandlerInterface: mysqlDBHandler,
	}

	service := &recordService.RecordCommandService{
		RecordCommandRepositoryInterface: &recordRepository.RecordCommandRepositoryCircuitBreaker{
			RecordCommandRepositoryInterface: repository,
		},
	}

	return service
}

func (k *kernel) recordQueryServiceContainer() *recordService.RecordQueryService {
	repository := &recordRepository.RecordQueryRepository{
		MySQLDBHandlerInterface: mysqlDBHandler,
	}

	service := &recordService.RecordQueryService{
		RecordQueryRepositoryInterface: &recordRepository.RecordQueryRepositoryCircuitBreaker{
			RecordQueryRepositoryInterface: repository,
		},
	}

	return service
}

func registerHandlers() {
	var err error

	// connect to database
	mysqlDBHandler = &mysql.MySQLDBHandler{}
	err = mysqlDBHandler.Connect(types.ConnectionParams{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBDatabase: os.Getenv("DB_DATABASE"),
		DBUsername: os.Getenv("DB_USERNAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatalf("[SERVER] mysql database is not responding: %v", err)
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
