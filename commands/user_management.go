package commands

import (
	"Go-Shell/models"
	"errors"
	"fmt"
)

func Login(args []string, currentUser **models.User) error {
	if len(args) != 2 {
		return errors.New("usage: login <username> <password>")
	}
	username, password := args[0], args[1]
	user := models.FindUser(username)
	if user == nil {
		return errors.New("user not found")
	}
	if user.Password != password {
		return errors.New("incorrect password")
	}
	*currentUser = user
	fmt.Printf("%s:$ \n", username)
	return nil
}

func AddUser(args []string) error {
	if len(args) != 2 {
		return errors.New("usage: adduser <username> <password>")
	}
	username, password := args[0], args[1]
	if models.FindUser(username) != nil {
		return errors.New("duplicate user exists with this username")
	}
	models.CreateUser(username, password)
	fmt.Println("user created successfully")
	return nil
}

func Logout(currentUser **models.User) error {
	*currentUser = models.NewUser("guest", "")
	return nil
}
