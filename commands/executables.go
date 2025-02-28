package commands

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

func ExecSystemCommand(args []string) error {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		var exitErr *exec.Error
		if errors.As(err, &exitErr) && errors.Is(exitErr.Err, exec.ErrNotFound) {
			fmt.Printf("%s: command not found\n", args[0])
			return nil
		}
		return err
	}
	return nil
}
