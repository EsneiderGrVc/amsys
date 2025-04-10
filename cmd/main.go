package main

import (
	"fmt"
	"log"
	"net/http"

	"amsys/configs"
	"amsys/internal/api"
	"amsys/internal/api/routes"
	"amsys/internal/web"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// API routes
	configs.BootstrapDatabase()
	apiRouter := r.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/users", api.GetUsers).Methods("GET")

	routes.RegisterRoutes(apiRouter)

	// // Web routes
	r.HandleFunc("/", web.HomeHandler).Methods("GET")
	r.HandleFunc("/home", web.HomeHandler).Methods("GET")
	r.HandleFunc("/users", web.UsersHandler).Methods("GET")
	r.HandleFunc("/paginate", web.TableHandler).Methods("GET")

	// // Serve static files (CSS, JS, images)
	// fs := http.FileServer(http.Dir("internal/static"))
	// r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	// Start the server
	fmt.Println("Server running on http://127.0.0.1:3001")
	log.Fatal(http.ListenAndServe("0.0.0.0:3001", r))
}
