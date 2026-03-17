package router

import (
	"github.com/gorilla/mux"
	"github.com/sh4shv4t/mongoapi/controller"
)

func Router(r *mux.Router) {
	r.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	r.HandleFunc("/api/movie/{id}", controller.GetOneMovie).Methods("GET")
	r.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	r.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	r.HandleFunc("/api/movie/{id}", controller.DeleteOneMovie).Methods("DELETE")
	r.HandleFunc("/api/deleteallmovie", controller.DeleteAllMovies).Methods("DELETE")
}
