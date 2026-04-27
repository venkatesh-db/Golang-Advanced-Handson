package main

import (
	"errors"
	"fmt"
)

// domain model
type User struct {
	ID   int
	Name string
}

// repository contract
type UserRepository interface {
	GetByID(id int) (User, error)
}

// concrete implementation
type InMemoryUserRepository struct {
	data map[int]User
}

func (r *InMemoryUserRepository) GetByID(id int) (User, error) {
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

func (s *UserService) GetUserName(id int) (string, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return "", err
	}
	return user.Name, nil
}

func main() {
	repo := &InMemoryUserRepository{
		data: map[int]User{
			1: {ID: 1, Name: "Ken thomson"},
		},
	}

	service := NewUserService(repo)

	name, err := service.GetUserName(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("user name:", name)
}
