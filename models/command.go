package models

import "time"

type Command struct {
	Name      string
	Count     int
	Timestamp time.Time
}

func NewCommand(name string) *Command {
	return &Command{
		Name:      name,
		Count:     1,
		Timestamp: time.Now(),
	}
}
