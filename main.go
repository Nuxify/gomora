package main

import (
	"log"
	"net/http"
)

func main() {

	err := http.ListenAndServe(":8080", ChiRouter().InitRouter())
	if err != nil {
		log.Println("Listen and Serve:", err)
	}
}
