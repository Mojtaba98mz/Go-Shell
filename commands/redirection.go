package commands

import (
	"fmt"
	"os"
)

func HandleRedirections(args *[]string) (origStdout, origStderr *os.File, stdoutFile, stderrFile *os.File, err error) {
	origStdout = os.Stdout
	origStderr = os.Stderr

	for i := 1; i < len(*args); i++ {
		switch (*args)[i] {
		case ">", "1>":
			if i+1 < len(*args) {
				stdoutFile, err = os.Create((*args)[i+1])
				if err != nil {
					return nil, nil, nil, nil, err
				}
				os.Stdout = stdoutFile
				*args = append((*args)[:i], (*args)[i+2:]...)
				i--
			} else {
				return nil, nil, nil, nil, fmt.Errorf("syntax error: expected file after %s", (*args)[i])
			}

		case ">>", "1>>":
			if i+1 < len(*args) {
				stdoutFile, err = os.OpenFile((*args)[i+1], os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					return nil, nil, nil, nil, err
				}
				os.Stdout = stdoutFile
				*args = append((*args)[:i], (*args)[i+2:]...)
				i--
			} else {
				return nil, nil, nil, nil, fmt.Errorf("syntax error: expected file after %s", (*args)[i])
			}

		case "2>":
			if i+1 < len(*args) {
				stderrFile, err = os.Create((*args)[i+1])
				if err != nil {
					return nil, nil, nil, nil, err
				}
				os.Stderr = stderrFile
				*args = append((*args)[:i], (*args)[i+2:]...)
				i--
			} else {
				return nil, nil, nil, nil, fmt.Errorf("syntax error: expected file after %s", (*args)[i])
			}

		case "2>>":
			if i+1 < len(*args) {
				stderrFile, err = os.OpenFile((*args)[i+1], os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					return nil, nil, nil, nil, err
				}
				os.Stderr = stderrFile
				*args = append((*args)[:i], (*args)[i+2:]...)
				i--
			} else {
				return nil, nil, nil, nil, fmt.Errorf("syntax error: expected file after %s", (*args)[i])
			}
		}
	}

	return origStdout, origStderr, stdoutFile, stderrFile, nil
}
