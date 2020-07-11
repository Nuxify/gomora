package grpc

import (
	"fmt"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"

	"gomora/interfaces"
	tenantGRPCPB "gomora/module/tenant/interfaces/http/grpc/pb"
)

// GRPCServerInterface holds the implementable method for the grpc server interface
type GRPCServerInterface interface {
	Serve(port int)
}

type server struct{}

var (
	m          *server
	serverOnce sync.Once
)

func (s *server) Serve(port int) {
	// create net listener
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("[SERVER] gRPC server failed %v", err)
	}

	// create grpc server
	grpcServer := grpc.NewServer()

	tenantServer := interfaces.ServiceContainer().RegisterTenantGRPCQueryController()
	tenantGRPCPB.RegisterTenantServiceServer(grpcServer, &tenantServer)

	log.Printf("[SERVER] gRPC server running on :%d", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("[SERVER] gRPC server failed %v", err)
	}
}

func registerHandlers() {}

// GRPCServer export instantiated grpc server once
func GRPCServer() GRPCServerInterface {
	if m == nil {
		serverOnce.Do(func() {
			// register http handlers
			registerHandlers()

			m = &server{}
		})
	}
	return m
}
