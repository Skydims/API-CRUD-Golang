// main.go
package main

import (
	"log"
	"net/http"

	"todo-app/database"
	"todo-app/routes"

	"github.com/gorilla/mux"
)

func main() {
	database.InitDB() // Inisialisasi database

	r := mux.NewRouter()
	routes.RegisterTaskRoutes(r)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", r)
}
