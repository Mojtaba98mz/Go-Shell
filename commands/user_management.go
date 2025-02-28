package commands

import (
	"Go-Shell/models"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

// to handle logout from one user and login as a previous user using stack
var userStack []*models.User

func Login(args []string, currentUser **models.User) error {
	if len(args) != 2 {
		return errors.New("usage: login <username> <password>")
	}
	username, password := args[0], args[1]
	user, err := models.FindUser(username)
	if err != nil {
		return errors.New("user not found")
	}
	if user.Password != password {
		return errors.New("incorrect password")
	}
	userStack = append(userStack, *currentUser)
	*currentUser = user
	return nil
}

func AddUser(args []string) error {
	if len(args) != 2 {
		return errors.New("usage: adduser <username> <password>")
	}
	username, password := args[0], args[1]
	user, err := models.FindUser(username)
	if user != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("duplicate user exists with this username")
	}
	models.NewUser(username, password)
	fmt.Println("user created successfully")
	return nil
}

func Logout(currentUser **models.User) error {
	if len(userStack) == 0 {
		*currentUser, _ = models.FindUser("guest")
		return nil
	}
	*currentUser = userStack[len(userStack)-1]
	userStack = userStack[:len(userStack)-1]
	return nil
}
