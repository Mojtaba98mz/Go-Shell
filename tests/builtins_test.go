package tests

import (
	"Go-Shell/commands"
	"Go-Shell/models"
	"os"
	"testing"
)

func TestExit(t *testing.T) {
	// Test exit with too many arguments
	t.Run("exit with too many arguments", func(t *testing.T) {
		err := commands.Exit([]string{"1", "2"})
		if err == nil || err.Error() != "too many arguments" {
			t.Errorf("Exit() error = %v, expected 'too many arguments'", err)
		}
	})
}

func TestEcho(t *testing.T) {
	// Test echo with no arguments
	t.Run("echo with no arguments", func(t *testing.T) {
		err := commands.Echo([]string{})
		if err != nil {
			t.Errorf("Echo() error = %v, expected nil", err)
		}
	})

	// Test echo with arguments
	t.Run("echo with arguments", func(t *testing.T) {
		err := commands.Echo([]string{"hello", "world"})
		if err != nil {
			t.Errorf("Echo() error = %v, expected nil", err)
		}
	})

	// Test echo with environment variable
	t.Run("echo with environment variable", func(t *testing.T) {
		os.Setenv("TEST_VAR", "value")
		err := commands.Echo([]string{"$TEST_VAR"})
		if err != nil {
			t.Errorf("Echo() error = %v, expected nil", err)
		}
	})
}

func TestCat(t *testing.T) {
	// Create a temporary file
	tmpFile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	// Write data to the temporary file
	data := "test data"
	if _, err := tmpFile.WriteString(data); err != nil {
		t.Fatal(err)
	}

	// Test cat with a valid file
	t.Run("cat with valid file", func(t *testing.T) {
		err := commands.Cat([]string{tmpFile.Name()})
		if err != nil {
			t.Errorf("Cat() error = %v, expected nil", err)
		}
	})

	// Test cat with no arguments
	t.Run("cat with no arguments", func(t *testing.T) {
		err := commands.Cat([]string{})
		if err == nil || err.Error() != "usage: cat <filename>" {
			t.Errorf("Cat() error = %v, expected 'usage: cat <filename>'", err)
		}
	})

	// Test cat with a non-existent file
	t.Run("cat with non-existent file", func(t *testing.T) {
		err := commands.Cat([]string{"nonexistent.txt"})
		if err == nil {
			t.Errorf("Cat() error = %v, expected file not found error", err)
		}
	})
}

func TestLs(t *testing.T) {
	// Create a temporary directory
	tmpDir, err := os.MkdirTemp("", "testdir")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Create a file in the temporary directory
	_, err = os.CreateTemp(tmpDir, "testfile")
	if err != nil {
		t.Fatal(err)
	}

	// Test ls with a valid directory
	t.Run("ls with valid directory", func(t *testing.T) {
		err := commands.Ls([]string{tmpDir})
		if err != nil {
			t.Errorf("Ls() error = %v, expected nil", err)
		}
	})

	// Test ls with no arguments
	t.Run("ls with no arguments", func(t *testing.T) {
		err := commands.Ls([]string{})
		if err != nil {
			t.Errorf("Ls() error = %v, expected nil", err)
		}
	})

	// Test ls with a non-existent directory
	t.Run("ls with non-existent directory", func(t *testing.T) {
		err := commands.Ls([]string{"nonexistent"})
		if err == nil {
			t.Errorf("Ls() error = %v, expected directory not found error", err)
		}
	})
}

func TestLl(t *testing.T) {
	// Create a temporary directory
	tmpDir, err := os.MkdirTemp("", "testdir")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Create a file in the temporary directory
	_, err = os.CreateTemp(tmpDir, "testfile")
	if err != nil {
		t.Fatal(err)
	}

	// Test ll with a valid directory
	t.Run("ll with valid directory", func(t *testing.T) {
		err := commands.Ll([]string{tmpDir})
		if err != nil {
			t.Errorf("Ll() error = %v, expected nil", err)
		}
	})

	// Test ll with no arguments
	t.Run("ll with no arguments", func(t *testing.T) {
		err := commands.Ll([]string{})
		if err != nil {
			t.Errorf("Ll() error = %v, expected nil", err)
		}
	})

	// Test ll with a non-existent directory
	t.Run("ll with non-existent directory", func(t *testing.T) {
		err := commands.Ll([]string{"nonexistent"})
		if err == nil {
			t.Errorf("Ll() error = %v, expected directory not found error", err)
		}
	})
}

func TestType(t *testing.T) {
	// Test type with a builtin command
	t.Run("type with builtin command", func(t *testing.T) {
		err := commands.Type([]string{"echo"})
		if err != nil {
			t.Errorf("Type() error = %v, expected nil", err)
		}
	})

	// Test type with an external command
	t.Run("type with external command", func(t *testing.T) {
		err := commands.Type([]string{"ls"})
		if err != nil {
			t.Errorf("Type() error = %v, expected nil", err)
		}
	})

	// Test type with a non-existent command
	t.Run("type with non-existent command", func(t *testing.T) {
		err := commands.Type([]string{"nonexistent"})
		if err == nil {
			t.Errorf("Type() error = %v, expected command not found error", err)
		}
	})

	// Test type with no arguments
	t.Run("type with no arguments", func(t *testing.T) {
		err := commands.Type([]string{})
		if err == nil || err.Error() != "usage: type <command>" {
			t.Errorf("Type() error = %v, expected 'usage: type <command>'", err)
		}
	})
}

func TestPwd(t *testing.T) {
	// Test pwd
	t.Run("pwd", func(t *testing.T) {
		err := commands.Pwd()
		if err != nil {
			t.Errorf("Pwd() error = %v, expected nil", err)
		}
	})
}

func TestCd(t *testing.T) {
	// Create a temporary directory
	tmpDir, err := os.MkdirTemp("", "testdir")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Test cd with a valid directory
	t.Run("cd with valid directory", func(t *testing.T) {
		err := commands.Cd([]string{tmpDir})
		if err != nil {
			t.Errorf("Cd() error = %v, expected nil", err)
		}
	})

	// Test cd with no arguments
	t.Run("cd with no arguments", func(t *testing.T) {
		err := commands.Cd([]string{})
		if err == nil || err.Error() != "usage: cd <directory>" {
			t.Errorf("Cd() error = %v, expected 'usage: cd <directory>'", err)
		}
	})

	// Test cd with a non-existent directory
	t.Run("cd with non-existent directory", func(t *testing.T) {
		err := commands.Cd([]string{"nonexistent"})
		if err == nil {
			t.Errorf("Cd() error = %v, expected directory not found error", err)
		}
	})
}

func TestClean(t *testing.T) {
	// Test clean
	t.Run("clean", func(t *testing.T) {
		user := &models.User{}
		err := commands.Clean(user)
		if err != nil {
			t.Errorf("Clean() error = %v, expected nil", err)
		}
	})
}
