package tests

import (
	"Go-Shell/commands"
	"Go-Shell/models"
	"testing"
)

func TestNewUser(t *testing.T) {
	user := models.NewUser("testuser", "password")
	if user == nil {
		t.Errorf("Expected user not to be nil")
	}
	if user.Username != "testuser" {
		t.Errorf("Expected username 'testuser', got %s", user.Username)
	}
}

func TestFindUser(t *testing.T) {
	models.NewUser("testuser", "password")

	user, err := models.FindUser("testuser")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if user.Username != "testuser" {
		t.Errorf("Expected username 'testuser', got %s", user.Username)
	}
}

func TestAddCommand(t *testing.T) {
	cleanUp()
	currentUser, _ := models.FindUser("guest")
	models.NewUser("testuser", "password")
	args := []string{"testuser", "password"}
	_ = commands.Login(args, &currentUser)
	currentUser.AddCommand("ls")

	history := currentUser.GetHistory()
	if len(history) != 1 {
		t.Errorf("Expected history length 1, got %d", len(history))
	}
	if history[0].Name != "ls" {
		t.Errorf("Expected command 'ls', got %s", history[0].Name)
	}
}

func TestClearHistory(t *testing.T) {
	user := models.NewUser("testuser", "password")
	user.AddCommand("ls")
	user.ClearHistory()

	history := user.GetHistory()
	if len(history) != 0 {
		t.Errorf("Expected history length 0, got %d", len(history))
	}
}

func TestDeleteGuestUser(t *testing.T) {
	models.NewUser("guest", "password")
	models.DeleteUser("guest")

	_, err := models.FindUser("guest")
	if err == nil {
		t.Errorf("Expected error when finding deleted guest user")
	}
}

func cleanUp() {
	models.DeleteUser("testuser")
}
