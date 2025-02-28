package commands

import (
	"Go-Shell/models"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var builtins = []string{
	"exit",
	"echo",
	"cat",
	"ls",
	"ll",
	"type",
	"pwd",
	"cd",
	"login",
	"adduser",
	"logout",
	"history",
	"clean",
}

func CommandNames() []string {
	return builtins
}

func Exit(args []string) error {
	if len(args) > 1 {
		return errors.New("too many arguments")
	}
	code := 0
	if len(args) == 1 {
		code, _ = strconv.Atoi(args[0])
	}
	fmt.Printf("exit status %d\n", code)
	os.Exit(code)
	return nil
}

func Echo(args []string) error {
	for _, arg := range args {
		if strings.HasPrefix(arg, "$") {
			arg = os.Getenv(arg[1:])
		}
		fmt.Print(arg, " ")
	}
	fmt.Println()
	return nil
}

func Cat(args []string) error {
	if len(args) != 1 {
		return errors.New("usage: cat <filename>")
	}
	data, err := os.ReadFile(args[0])
	if err != nil {
		return err
	}
	fmt.Print(string(data))
	return nil
}

func Ls(args []string) error {
	path := "."
	if len(args) > 0 {
		path = args[0]
	}
	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
	return nil
}

func Ll(args []string) error {
	path := "."
	if len(args) > 0 {
		path = args[0]
	}
	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	for _, file := range files {
		info, err := file.Info()
		if err != nil {
			return err
		}
		fmt.Printf("%s\t%d\t%s\n", info.Mode(), info.Size(), file.Name())
	}
	return nil
}

func Type(args []string) error {
	if len(args) != 1 {
		return errors.New("usage: type <command>")
	}
	cmd := args[0]
	if isBuiltin(cmd) {
		fmt.Printf("%s is a shell builtin\n", cmd)
	} else {
		path, err := exec.LookPath(cmd)
		if err != nil {
			return err
		}
		fmt.Printf("%s is %s\n", cmd, path)
	}
	return nil
}

func Pwd() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Println(dir)
	return nil
}

func Cd(args []string) error {
	if len(args) != 1 {
		return errors.New("usage: cd <directory>")
	}
	return os.Chdir(args[0])
}

func Clean(user *models.User) error {
	user.ClearHistory()
	return nil
}

func isBuiltin(cmd string) bool {
	for _, b := range builtins {
		if b == cmd {
			return true
		}
	}
	return false
}
