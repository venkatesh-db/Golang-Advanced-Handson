
package main 

type UserService interface {

	 GetUser(id int ) string 
	 CreateUser(name string) error
}

