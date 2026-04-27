package main

import (
	"errors"
	"fmt"
	"reflect"
)

// =======================
// 🔷 GENERIC API RESPONSE
// =======================
type ApiResponse[T any] struct {
	Status  string
	Data    T
	Message string
}

// =======================
// 🔷 USER MODEL (REAL WORLD)
// =======================
type User struct {
	ID    int    `required:"true"`
	Name  string `required:"true"`
	Email string `required:"true"`
	Age   int    `required:"true"`
}

// =======================
// 🔷 GENERIC REPOSITORY
// =======================
type Repository[T any] struct {
	store map[int]T
}

func NewRepository[T any]() *Repository[T] {
	return &Repository[T]{
		store: make(map[int]T),
	}
}

func (r *Repository[T]) Save(id int, data T) {
	r.store[id] = data
}

func (r *Repository[T]) Get(id int) (T, error) {
	val, ok := r.store[id]
	if !ok {
		var zero T
		return zero, errors.New("not found")
	}
	return val, nil
}

// =======================
// 🔷 REFLECTION VALIDATOR
// =======================
func ValidateStruct(input interface{}) error {
	val := reflect.ValueOf(input)

	// Handle pointer
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	typ := reflect.TypeOf(input)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	for i := 0; i < val.NumField(); i++ {
		fieldValue := val.Field(i)
		fieldType := typ.Field(i)

		requiredTag := fieldType.Tag.Get("required")

		if requiredTag == "true" && fieldValue.IsZero() {
			return fmt.Errorf("field %s is required", fieldType.Name)
		}
	}

	return nil
}

// =======================
// 🔷 SERVICE LAYER
// =======================
type UserService struct {
	repo *Repository[User]
}

func NewUserService() *UserService {
	return &UserService{
		repo: NewRepository[User](),
	}
}

func (s *UserService) CreateUser(u User) ApiResponse[User] {

	// 🔥 Reflection Validation
	if err := ValidateStruct(u); err != nil {
		return ApiResponse[User]{
			Status:  "error",
			Message: err.Error(),
		}
	}

	// Save using Generics Repo
	s.repo.Save(u.ID, u)

	return ApiResponse[User]{
		Status: "success",
		Data:   u,
	}
}

func (s *UserService) GetUser(id int) ApiResponse[User] {
	user, err := s.repo.Get(id)

	if err != nil {
		return ApiResponse[User]{
			Status:  "error",
			Message: "User not found",
		}
	}

	return ApiResponse[User]{
		Status: "success",
		Data:   user,
	}
}

// =======================
// 🔷 MAIN FUNCTION
// =======================
func main() {

	service := NewUserService()

	// ✅ VALID USER
	user := User{
		ID:    1,
		Name:  "Venkatesh",
		Email: "venkat@gmail.com",
		Age:   39,
	}

	resp := service.CreateUser(user)
	fmt.Println("Create Response:", resp)

	//  INVALID USER (missing Name)
	badUser := User{
		ID:    2,
		Email: "bad@gmail.com",
	}

	resp2 := service.CreateUser(badUser)
	fmt.Println("Create Response:", resp2)

	//  FETCH USER
	getResp := service.GetUser(1)
	fmt.Println("Get Response:", getResp)
}