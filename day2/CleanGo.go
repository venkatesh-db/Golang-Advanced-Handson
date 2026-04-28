package main

import (
	"errors"
	"fmt"
)

// ---------------- DOMAIN ----------------

type User struct {
	ID   int
	Name string
}

// ---------------- INTERFACE ----------------

type UserRepository interface {
	Save(user User) error
	GetByID(id int) (User, error)
}

// ---------------- IMPLEMENTATION ----------------

type InMemoryUserRepo struct {
	data map[int]User
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{
		data: make(map[int]User),
	}
}

func (r *InMemoryUserRepo) Save(user User) error {
	if user.ID == 0 {
		return errors.New("invalid user ID")
	}
	r.data[user.ID] = user
	return nil
}

func (r *InMemoryUserRepo) GetByID(id int) (User, error) {
	user, ok := r.data[id]
	if !ok {
		return User{}, errors.New("user not found")
	}
	return user, nil
}

// ---------------- SERVICE ----------------

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(id int, name string) error {
	if name == "" {
		return errors.New("name cannot be empty")
	}

	user := User{
		ID:   id,
		Name: name,
	}

	return s.repo.Save(user)
}

func (s *UserService) GetUser(id int) (User, error) {
	return s.repo.GetByID(id)
}

// ---------------- APPLICATION ----------------

func main() {
	repo := NewInMemoryUserRepo()
	service := NewUserService(repo)

	// Create user
	err := service.CreateUser(1, "Venkatesh")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	// Fetch user
	user, err := service.GetUser(1)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("User: %+v\n", user)
}
