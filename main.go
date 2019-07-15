/*
	Googol API Framework
	A minimalist web API framework using the service-go-pattern

	@Author: Karl Anthony B. Baluyot
*/

package main

import (
	"log"
	"net/http"
)

//* entry point of the googol framework
func main() {

	err := http.ListenAndServe(":8080", ChiRouter().InitRouter())
	if err != nil {
		log.Println("Listen and Serve:", err)
	}
}
