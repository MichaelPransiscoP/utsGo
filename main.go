package main

import (
	"fmt"
	"log"
	"net/http"
	"uts/controllers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Define your routes using controllers
	router.HandleFunc("/rooms", controllers.GetAllRooms).Methods("GET")
	router.HandleFunc("/rooms/{id}", controllers.GetDetailRoom).Methods("GET")
	router.HandleFunc("/rooms", controllers.InsertRoom).Methods("POST")
	router.HandleFunc("/rooms/{id}", controllers.LeaveRoom).Methods("DELETE")

	// Start the HTTP server
	port := 8888
	fmt.Printf("Server started on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
