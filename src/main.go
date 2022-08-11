package main

import (
	"fmt"
	"net/http"

	"talkyapi/router"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("starting up HTTP server...")
	fmt.Println("Started")

	r := mux.NewRouter()
	
	router.SetupRoutes(r)


	err := http.ListenAndServe(":8000", r)
	if err != nil {
			panic(err)
	}
}