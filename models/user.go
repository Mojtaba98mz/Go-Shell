package models

import (
	"Go-Shell/database"
	"time"
)

type User struct {
	ID       uint      `gorm:"primaryKey"`
	Username string    `gorm:"unique;not null"`
	Password string    `gorm:"not null"`
	History  []Command `gorm:"foreignKey:UserID"`
}

func NewUser(username, password string) *User {
	user := &User{Username: username, Password: password}
	database.GetDB().Create(user)
	return user
}

func FindUser(username string) (*User, error) {
	var user User
	result := database.GetDB().Preload("History").Where("username = ?", username).First(&user)
	return &user, result.Error
}

func (u *User) AddCommand(name string) {
	cmd := Command{Name: name, Count: 1, Timestamp: time.Now(), UserID: u.ID}
	u.History = append(u.History, cmd)
	database.GetDB().Save(u)
}

func (u *User) GetHistory() []Command {
	var commands []Command
	database.GetDB().Where("user_id = ?", u.ID).Find(&commands)
	return commands
}

func (u *User) ClearHistory() {
	database.GetDB().Where("user_id = ?", u.ID).Delete(&Command{})
	u.History = nil
}

func DeleteUser(username string) {
	var user User
	if err := database.GetDB().Where("username = ?", username).Preload("History").First(&user).Error; err == nil {
		database.GetDB().Delete(&user)
	}
}
