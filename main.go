package main

import (
	"Go-Shell/commands"
	"Go-Shell/models"
	"fmt"
	"github.com/peterh/liner"
	"os"
	"strings"
)

var currentUser *models.User

func main() {
	models.DeleteUser("guest")
	currentUser = models.NewUser("guest", "")
	line := liner.NewLiner()
	defer line.Close()

	// set Ctrl+c to abort
	line.SetCtrlCAborts(true)

	// basic commands code completion support
	line.SetCompleter(func(line string) (c []string) {
		for _, n := range commands.CommandNames() {
			if strings.HasPrefix(n, line) {
				c = append(c, n)
			}
		}
		return
	})

	fmt.Println("Welcome to the Go Shell! Type 'exit' to quit.")

	for {
		if input, err := line.Prompt(currentUser.Username + ":$ "); err == nil {
			if err := processInput(input); err != nil {
				fmt.Println("Error:", err)
			}
			line.AppendHistory(input)
		} else if err == liner.ErrPromptAborted {
			fmt.Println("Aborted")
		} else {
			break
		}
	}
}

func processInput(input string) error {
	args := parseArguments(input)
	if len(args) == 0 {
		return nil
	}

	command := args[0]
	currentUser.AddCommand(input)

	// Handle redirections
	origStdout, origStderr, err := commands.HandleRedirections(&args)
	if err != nil {
		return err
	}
	defer func() {
		os.Stdout = origStdout
		os.Stderr = origStderr
	}()

	switch command {
	case "exit":
		return commands.Exit(args[1:])
	case "echo":
		return commands.Echo(args[1:])
	case "cat":
		return commands.Cat(args[1:])
	case "ls":
		return commands.Ls(args[1:])
	case "ll":
		return commands.Ll(args[1:])
	case "type":
		return commands.Type(args[1:])
	case "pwd":
		return commands.Pwd()
	case "cd":
		return commands.Cd(args[1:])
	case "login":
		return commands.Login(args[1:], &currentUser)
	case "adduser":
		return commands.AddUser(args[1:])
	case "logout":
		return commands.Logout(&currentUser)
	case "history":
		return commands.History(args[1:], currentUser)
	case "clean":
		return commands.Clean(currentUser)
	default:
		return commands.ExecSystemCommand(args)
	}
}

func parseArguments(input string) []string {
	var args []string
	var current strings.Builder
	inQuotes := false
	escaped := false

	for _, r := range input {
		switch {
		case escaped:
			switch r {
			case '$', '`', '"', '\\', '\n':
				current.WriteRune(r)
			default:
				current.WriteRune('\\')
				current.WriteRune(r)
			}
			escaped = false
		case r == '\\':
			escaped = true
		case r == '"':
			inQuotes = !inQuotes
		case !inQuotes && (r == ' ' || r == '\t'):
			if current.Len() > 0 {
				args = append(args, current.String())
				current.Reset()
			}
		default:
			current.WriteRune(r)
		}
	}

	if current.Len() > 0 {
		args = append(args, current.String())
	}

	return args
}
