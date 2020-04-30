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
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func init() {
	err := godotenv.Load(fmt.Sprintf("%s/.env", basepath))
	if err != nil {
		panic(err)
	}
}

func main() {
	port := os.Getenv("API_URL_PORT")
	if len(port) == 0 {
		port = "8000" // default port is 8000 if not set
	}

	// run listeners concurrently here:

	// run http server
	log.Println("[SERVER] Listen and Serve: localhost:", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), ChiRouter().InitRouter())
	if err != nil {
		log.Println("Error:", err)
	}
}
