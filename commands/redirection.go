package commands

import (
	"fmt"
	"os"
)

func HandleRedirections(args *[]string) (origStdout, origStderr *os.File, err error) {
	origStdout = os.Stdout
	origStderr = os.Stderr

	var stdoutFile, stderrFile *os.File
	defer func() {
		if stdoutFile != nil {
			stdoutFile.Close()
		}
		if stderrFile != nil {
			stderrFile.Close()
		}
	}()

	for i := 1; i < len(*args); i++ {
		if (*args)[i] == ">" || (*args)[i] == "1>" {
			if i+1 < len(*args) {
				stdoutFile, err = os.Create((*args)[i+1])
				if err != nil {
					return nil, nil, err
				}
				*args = append((*args)[:i], (*args)[i+2:]...)
				i--
			} else {
				return nil, nil, fmt.Errorf("syntax error: expected file after %s", (*args)[i])
			}
		} else if (*args)[i] == ">>" || (*args)[i] == "1>>" {
			if i+1 < len(*args) {
				stdoutFile, err = os.OpenFile((*args)[i+1], os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					return nil, nil, err
				}
				*args = append((*args)[:i], (*args)[i+2:]...)
				i--
			} else {
				return nil, nil, fmt.Errorf("syntax error: expected file after %s", (*args)[i])
			}
		} else if (*args)[i] == "2>" {
			if i+1 < len(*args) {
				stderrFile, err = os.Create((*args)[i+1])
				if err != nil {
					return nil, nil, err
				}
				*args = append((*args)[:i], (*args)[i+2:]...)
				i--
			} else {
				return nil, nil, fmt.Errorf("syntax error: expected file after %s", (*args)[i])
			}
		} else if (*args)[i] == "2>>" {
			if i+1 < len(*args) {
				stderrFile, err = os.OpenFile((*args)[i+1], os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					return nil, nil, err
				}
				*args = append((*args)[:i], (*args)[i+2:]...)
				i--
			} else {
				return nil, nil, fmt.Errorf("syntax error: expected file after %s", (*args)[i])
			}
		}
	}

	if stdoutFile != nil {
		os.Stdout = stdoutFile
	}
	if stderrFile != nil {
		os.Stderr = stderrFile
	}

	return origStdout, origStderr, nil
}
