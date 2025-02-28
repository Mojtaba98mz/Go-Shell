# Go-Shell

![Go](https://img.shields.io/badge/Go-1.16-blue?logo=go)

## 📋 Table of Contents
1. [Project Overview](#overview)
2. [Core Features](#core-features)
3. [Project Structure](#project-structure)
4. [Minimum Requirements](#minimum-requirements)
5. [Build and Deploy](#build-and-deploy)
6. [Configuration](#configuration)
7. [Contact Us](#contact-us)
8. [License](#license)

## Overview
The **Go-Shell** is a command-line shell written in Go. It supports built-in commands, external command execution, command history, user management, input/output redirections, and more.
I also use liner to support up arrow for last history commands and Ctrl+c to abort a command

## Core Features
- **Built-in Commands:** Implementation of common shell commands.
- **Command Execution:** Ability to execute both built-in and external commands.
- **Command History:** Maintains a history of executed commands.
- **User Management:** Supports user login and logout functionality.
- **Input/Output Redirection:** Handles input/output redirections.
- **Utility Functions:** Provides various utility functions.
- **Database Support:** Connects and queries the database.

---

## Project Structure
```plaintext
shell-go/
├── commands/                       # Command executors and built-in commands
│   ├── builtins.go                 # Implementation of built-in commands
│   ├── executables.go              # Runs commands (built-in & external)
│   ├── history.go                  # Handles command history
│   ├── redirection.go              # Handle redirections
│   ├── user_management.go          # Manages users (login/logout)
├── model/                          # Internal utilities (helpers, logging, etc.)
│   ├── command.go                  # Handle commands data
│   ├── user.go                     # Handle users data
├── database/                       # Database-related code
│   ├── database.go                 # Connects and queries the database
├── tests/                          # Unit tests
│   ├── builtins_test.go            # Tests for built-in commands
│   ├── user_management_test.go     # Tests for the command executor
│   ├── user_test.go                # Tests for user management
├── main.go                         # Entry point of the shell
├── README.md                       # Documentation
├── go.mod                          # Go module file
├── go.sum                          # Dependency lock file
```

---

## Minimum Requirements
To run the application, ensure you have the following:
```plaintext
| Component | Version |
|-----------|---------|
| **Go**    | 1.16+   |
```

---

## Build and Deploy

### Step 1: Clone the Repository
```bash
git clone https://github.com/Mojtaba98mz/Go-Shell.git
cd Go-Shell
```

### Step 2: Build the Application
```bash
go build -o go-shell main.go
```

### Step 3: Run the Application
```bash
./go-shell
```

---

## Configuration
- **Database:** Ensure your database is running and accessible.
- **Environment Variables:**
    - `DB_USERNAME`: Database username.
    - `DB_PASSWORD`: Database password.

---

## 📞 Contact Us
For any questions or issues, feel free to reach out:

| Name              | Contact Info                                                           |
|-------------------|------------------------------------------------------------------------|
| Mojtaba Zamandi   | mojtabazamandi.mz@gmail.com                                            |
| GitHub            | [Mojtaba98mz](https://github.com/Mojtaba98mz)                          |
| GitHub Issues     | [Open an Issue](https://github.com/Mojtaba98mz/Go-Shell/issues)        |
| LinkedIn          | [Mojtaba Zamandi](https://linkedin.com/in/mojtaba-zamandi)             |

---

## License
This project is licensed under the MIT License - see the full text below.

MIT License

```
MIT License

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```