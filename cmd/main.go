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
)

func init() {
	// load our environmental variables.
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {
	grpcPort, err := strconv.Atoi(os.Getenv("API_URL_GRPC_PORT"))
	if err != nil {
		log.Fatalf("[SERVER] Invalid port")
	}
	if len(fmt.Sprintf("%d", grpcPort)) == 0 {
		grpcPort = 9090 // default grpcPort is 9090 if not set
	}

	// serve grpc server
	grpc.GRPCServer().Serve(grpcPort)
}
