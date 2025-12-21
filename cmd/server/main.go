package main

import (
	"github.com/gorilla/mux"

	"advanced-go-app/internal/config"
	"advanced-go-app/internal/middleware"
	"advanced-go-app/internal/server"
	"advanced-go-app/internal/user"
)

func main() {
	cfg := config.Load()

	repo := user.NewMemoryRepo()
	service := user.NewService(repo)
	handler := user.NewHandler(service)

	r := mux.NewRouter()

	admin := r.PathPrefix("/admin").Subrouter()
	admin.Use(middleware.Role("admin"))
	admin.HandleFunc("/users", handler.GetUsers).Methods("GET")

	srv := server.New(":"+cfg.Port, r)
	srv.Start()
}
