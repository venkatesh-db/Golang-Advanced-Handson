package main

// http://localhost:8080/user?id=1

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
)

// domain model
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// repository contract
type UserRepository interface {
	GetByID(ctx context.Context, id int) (User, error)
}

// repository implementation
type InMemoryUserRepository struct {
	data map[int]User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		data: map[int]User{
			1: {ID: 1, Name: "Venkatesh"},
			2: {ID: 2, Name: "Arjun"},
		},
	}
}

func (r *InMemoryUserRepository) GetByID(ctx context.Context, id int) (User, error) {
	user, ok := r.data[id]
	if !ok {
		return User{}, errors.New("user not found")
	}
	return user, nil
}

// service layer
type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUser(ctx context.Context, id int) (User, error) {
	return s.repo.GetByID(ctx, id)
}

// HTTP handler
type UserHandler struct {
	service *UserService
}

func NewUserHandler(service *UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUser(r.Context(), id)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// server setup
func main() {
	repo := NewInMemoryUserRepository()
	service := NewUserService(repo)
	handler := NewUserHandler(service)

	mux := http.NewServeMux()
	mux.HandleFunc("/user", handler.GetUser)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("server running on http://localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
