package main

import (
	"API_CRUD/models"
	"API_CRUD/routes"
	"log/slog"
	"net/http"
	"time"
)

func main(){
	if err := run(); err != nil{
		slog.Error("Failed to execute code", "error", err)
		return
	}
	slog.Info("All systems offline")
}

func run() error {
	db := &models.Application{
		Data: make(map[models.ID]models.User),
	}

	handler := routes.NewHandler(db)

	s := http.Server{
		ReadTimeout: 10 * time.Second,
		IdleTimeout: time.Minute,
		WriteTimeout: 10 * time.Second,
		Addr: ":8080",
		Handler: handler,
	}
	
	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}