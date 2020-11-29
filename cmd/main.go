/*
|--------------------------------------------------------------------------
| Main
|--------------------------------------------------------------------------
|
| This is the entry point for listeners of the project.
| You can create and run goroutines for event listeners below before the HTTP listener.
|
*/
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"

	"gomora/interfaces/http/grpc"
	"gomora/interfaces/http/rest"
)

func init() {
	// load our environmental variables.
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {
	// grpc port
	grpcPort, err := strconv.Atoi(os.Getenv("API_URL_GRPC_PORT"))
	if err != nil {
		log.Fatalf("[SERVER] Invalid port")
	}
	if len(fmt.Sprintf("%d", grpcPort)) == 0 {
		grpcPort = 9090 // default grpcPort is 9090 if not set
	}

	// rest port
	restPort, err := strconv.Atoi(os.Getenv("API_URL_REST_PORT"))
	if err != nil {
		log.Fatalf("[SERVER] Invalid port")
	}
	if len(fmt.Sprintf("%d", restPort)) == 0 {
		restPort = 8000 // default grpcPort is 8000 if not set
	}

	// serve rest server
	go rest.ChiRouter().Serve(restPort)

	// serve grpc server
	grpc.GRPCServer().Serve(grpcPort)
}
