package models

var users = map[string]*User{}

type User struct {
	Username string
	Password string
	History  []*Command
}

func NewUser(username, password string) *User {
	return &User{
		Username: username,
		Password: password,
		History:  []*Command{},
	}
}

func FindUser(username string) *User {
	return users[username]
}

func CreateUser(username, password string) {
	users[username] = NewUser(username, password)
}

func (u *User) AddCommand(cmd *Command) {
	u.History = append(u.History, cmd)
}

func (u *User) GetHistory() []*Command {
	return u.History
}

func (u *User) ClearHistory() {
	u.History = []*Command{}
}
