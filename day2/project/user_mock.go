
package main

// Code generated automatically. DO NOT EDIT.

type UserServiceMock struct{}


func (m *UserServiceMock) GetUser(id int) string {
	// mock logic
	return "mock-user"
}

func (m *UserServiceMock) CreateUser(name string) error {
	// mock logic
	return nil
}
