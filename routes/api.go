package routes

import (
	"encoding/json"
	"log/slog"
	"API_CRUD/models"
	"API_CRUD/services"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHandler(db *models.Application) http.Handler{
	router := chi.NewMux()

	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)

	router.Post("/api/users", handlePost(db))
	router.Get("/api/users", handleGet(db))
	router.Get("/api/users/{id}", handleGetID(db))
	router.Delete("/api/users/{id}", handleDelete(db))
	router.Put("/api/users/{id}", handlePut(db))

	return router
}

type PostBody struct{
	URL string `json:"url"`
}

type Response struct{
	Error string `json:"error,omitempty"`
	Data any `json:"data,omitempty"`
}

func sendJSON(w http.ResponseWriter, resp Response, status int){
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(resp)
	if err != nil{
		slog.Error("failed to marshal json data", "error", err)
		sendJSON(
			w,
			Response{Error: "something went wrong"},
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil{
		slog.Error("Failed to write response to client", "error", err)
		return
	}
}

func handlePost(db *models.Application) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		var user models.User

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil{
			sendJSON(
				w,
				Response{Error: "invalid format"},
				http.StatusInternalServerError,
			)
			return
		}

		insertedUser, err := services.Insert(db, user)
		if err != nil{
			sendJSON(
				w,
				Response{Error: "failed to insert"},
				http.StatusInternalServerError,
			)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(insertedUser)
	}
}

func handleGet(db *models.Application) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){
		users := services.FindAll(db)

		if users == nil {
			sendJSON(
				w,
				Response{Error: "the database is empty"},
				http.StatusInternalServerError,
			)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}

func handleGetID(db *models.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		user, err := services.FindByID(db, id)
		if err != nil {
			if err.Error() == "user not found" {
				sendJSON(
					w,
					Response{Error: "The user with the specified ID does not exist"},
					http.StatusNotFound,
				)
				return
			} else {
				sendJSON(
					w,
					Response{Error: "The user information could not be retrieved"},
					http.StatusInternalServerError,
				)
				return
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}

func handleDelete(db *models.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		user, err := services.Delete(db, id)
		if err != nil {
			if err.Error() == "user not found" {
				sendJSON(
					w,
					Response{Error: "The user with the specified ID does not exist"},
					http.StatusNotFound,
				)
				return
			} else {
				sendJSON(
					w,
					Response{Error: "The user information could not be retrieved"},
					http.StatusInternalServerError,
				)
				return
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}

func handlePut(db *models.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		id := chi.URLParam(r, "id")
		
		var inputUser models.User
		if err := json.NewDecoder(r.Body).Decode(&inputUser); err != nil {
			sendJSON(
				w,
				Response{Error: "Invalid resquest body"},
				http.StatusBadRequest,
			)
			return
		}

		if inputUser.FirstName == "" || inputUser.Biography == "" {
			sendJSON(
				w,
				Response{Error: "Please provide name and bio for the user"},
				http.StatusBadRequest,
			)
			return
		}

		updatedUser, err := services.Update(db, id, inputUser)
		if err != nil {
			if err.Error() == "user not found"{
				sendJSON(
					w,
					Response{Error: "The user with the specified ID does not exist"},
					http.StatusNotFound,
				)
				return
			} else {
				sendJSON(
					w,
					Response{Error: "The user information could not be modified"},
					http.StatusInternalServerError,
				)
				return
			}
		}

		sendJSON(
			w,
			Response{Data: updatedUser},
			http.StatusOK,
		)
	}
}