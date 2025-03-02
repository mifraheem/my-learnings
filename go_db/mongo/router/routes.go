package router

import (
	"mongo/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/movies", controller.AllMovies).Methods("GET")
	router.HandleFunc("/api/v1/movies", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/v1/movies", controller.DeleteAllMovies).Methods("DELETE")
	router.HandleFunc("/api/v1/movies/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/v1/movies/{id}", controller.DeleteOneMovie).Methods("DELETE")

	return router
}
