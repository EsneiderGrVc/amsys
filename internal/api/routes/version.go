package routes

import (
	api_handlers "amsys/internal/api/handlers"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	versionRouter := r.PathPrefix("/version").Subrouter()
	versionRouter.HandleFunc("/", api_handlers.GetProjectHandler).Methods("GET")
	versionRouter.HandleFunc("/", api_handlers.CreateVersionHandler).Methods("POST")
	versionRouter.HandleFunc("/ga/", api_handlers.CreateVersionFromGithubHandler).Methods("POST")
}
