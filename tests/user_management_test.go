package tests

import (
	"Go-Shell/commands"
	"Go-Shell/models"
	"errors"
	"testing"
)

func TestLogin(t *testing.T) {
	// Setup initial users in the database
	models.NewUser("user1", "password1")
	models.NewUser("user2", "password2")

	tests := []struct {
		args          []string
		expectedError error
		setupUsers    []string // List of users to add before testing
		expectedUser  string
	}{
		{
			args:          []string{"user1", "password1"},
			expectedError: nil,
			setupUsers:    []string{"user1", "user2"},
			expectedUser:  "user1",
		},
		{
			args:          []string{"user1", "wrongpassword"},
			expectedError: errors.New("incorrect password"),
			setupUsers:    []string{"user1", "user2"},
			expectedUser:  "user1",
		},
		{
			args:          []string{"nonexistent", "password"},
			expectedError: errors.New("user not found"),
			setupUsers:    []string{"user1", "user2"},
			expectedUser:  "user1",
		},
	}

	for _, tt := range tests {
		// Setup users in the database before each test
		for _, username := range tt.setupUsers {
			models.NewUser(username, "password")
		}

		currentUser := &models.User{}
		err := commands.Login(tt.args, &currentUser)

		if err != nil && err.Error() != tt.expectedError.Error() {
			t.Errorf("expected error: %v, got: %v", tt.expectedError, err)
		}

		// Check if user state is as expected
		if err == nil && currentUser.Username != tt.expectedUser {
			t.Errorf("expected logged in user: %v, got: %v", tt.expectedUser, currentUser.Username)
		}
	}
}

func TestAddUser(t *testing.T) {
	models.DeleteUser("john")
	tests := []struct {
		name    string
		args    []string
		wantErr bool
		errMsg  string
	}{
		{
			name:    "successful user creation",
			args:    []string{"john", "password123"},
			wantErr: false,
			errMsg:  "",
		},
		{
			name:    "duplicate user",
			args:    []string{"john", "password123"},
			wantErr: true,
			errMsg:  "duplicate user exists with this username",
		},
		{
			name:    "insufficient arguments",
			args:    []string{"john"},
			wantErr: true,
			errMsg:  "usage: adduser <username> <password>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := commands.AddUser(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil && err.Error() != tt.errMsg {
				t.Errorf("AddUser() error message = %v, want %v", err.Error(), tt.errMsg)
			}
		})
	}
}
