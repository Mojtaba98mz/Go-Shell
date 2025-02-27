# Bill Management System

![Bill Management System](https://img.shields.io/badge/Spring-Boot-green?logo=springboot) ![Docker](https://img.shields.io/badge/Docker-Enabled-blue?logo=docker) ![Java](https://img.shields.io/badge/Java-17-orange?logo=java)

## 📋 Table of Contents
1. [Project Overview](#overview)
2. [Swagger UI](#swagger-ui)
3. [CI CD Pipeline](#ci-cd-pipeline)
4. [Core Features](#core-features)
5. [Project Structure](#project-structure)
6. [Entities](#key-entities)
7. [Minimum Requirements](#minimum-requirements)
8. [Build and Deploy](#build-and-deploy)
9. [API Endpoints](#api-endpoints)
10. [Configuration](#configuration)
11. [Contact Us](#-contact-us)
12. [License](#license)

## Overview
The **Bill Management System** is a Java Spring-based application that simplifies expense sharing among groups. It tracks user contributions, computes balances, and determines how much money members owe each other to settle debts.

## Swagger UI

To access the Swagger UI for the API documentation, navigate to:

- **Swagger UI URL**: `http://185.204.197.57:8080/swagger-ui/index.html`

The Swagger UI will give you a user-friendly interface to interact with your API endpoints.

---

## CI CD Pipeline

This project is deployed using a CI/CD pipeline to automate build, test, and deployment processes. Below is an overview of the pipeline configuration.

### CI/CD Pipeline Flow:

1. **Build Job**:
    - Triggered on each push to the `master` branch.
    - The pipeline checks out the code from the repository, sets up Java 17, and runs `mvn clean package` to build the project and run tests.

2. **Deploy Job**:
    - After the build job is successfully completed, the deployment job is triggered.
    - The job sets up SSH keys and removes the previous `app.jar` from the server.
    - The new `app.jar` is copied to the server and the deployment script (`deploy.sh`) is executed.

### Pipeline Configuration:
You can view the full pipeline configuration in the `.github/workflows/deploy.yml` file.

### Core Features
- **User and Group Management:** Users can create and join groups with descriptive titles.
- **Member Tracking:** Groups can have multiple members, each identified by their name.
- **Expense Logging:** Members can pay for group expenses, recorded in the `Bill` entity with an amount field.
- **Expense Balancing:** Automatically calculates and displays how much money each member owes or is owed to balance expenses.
- **JWT Authentication:** Ensures secure access to APIs with token-based authentication.

---

## Project Structure
```plaintext
shell-go/
│── cmd/                     # Command executors and built-in commands
│   ├── builtins.go          # Implementation of built-in commands
│   ├── executor.go          # Runs commands (built-in & external)
│   ├── history.go           # Handles command history
│   ├── user.go              # Manages users (login/logout)
│── internal/                # Internal utilities (helpers, logging, etc.)
│   ├── parser.go            # Parses user input
│   ├── redirect.go          # Handles input/output redirections
│   ├── utils.go             # Utility functions
│── db/                      # Database-related code
│   ├── db.go                # Connects and queries the database
│── config/                  # Configurations
│   ├── config.go            # Stores configurations (e.g., paths, settings)
│── tests/                   # Unit tests
│   ├── builtins_test.go     # Tests for built-in commands
│   ├── executor_test.go     # Tests for the command executor
│   ├── parser_test.go       # Tests for input parsing
│── main.go                  # Entry point of the shell
│── README.md                # Documentation
│── go.mod                   # Go module file
│── go.sum                   # Dependency lock file
    
```

### Key Entities
| Entity   | Description                                                                           |
|----------|---------------------------------------------------------------------------------------|
| **User** | Represents a system user who can create and join groups.                              |
| **Role** | Defines user roles for access control.                                                |
| **Group**| Represents a group with a title, managed by users.                                    |
| **Member**| Represents individuals in a group who share expenses.                                |
| **Bill** | Tracks payments made by members, including the amount and associated group.           |

---

## Minimum Requirements
To run the application, ensure you have the following:

| Component   | Version |
|-------------|---------|
| **JDK**     | 17+     |
| **Spring Boot**| 3.0+ |
| **Docker**  | 20.10+  |
| **PostgreSQL** | 13+  |
| **Maven**   | 3.8+    |

---

## Build and Deploy

### Step 1: Clone the Repository
```bash
git clone https://github.com/Mojtaba98mz/BillManagement.git
cd BillManagement
```

### Step 2: Configure the Application
Edit the `application.properties` file in the `src/main/resources` directory:
```properties
spring.datasource.url=jdbc:postgresql://localhost:5432/your_db_name
spring.datasource.username=username
spring.datasource.password=password
spring.jpa.properties.hibernate.dialect=org.hibernate.dialect.PostgreSQLDialect
spring.jpa.hibernate.ddl-auto=update
spring.jpa.show-sql=true

jwtSecret: your_jwt_secret_key
tokenValidityInSeconds: 3600
```

### Step 3: Build the Application
```bash
./mvnw clean install
```

### Step 4: Run the Application
#### Using Docker:
1. Build the Docker image:
   ```bash
   docker build -t bill-management .
   ```
   or
   ```bash
   ./mvnw jib:dockerBuild
   ```

2. Run the container:
   ```bash
   docker run -p 8080:8080 bill-management
   ```

#### Using Maven:
```bash
./mvnw spring-boot:run
```
#### Using Docker Compose:
```bash
docker-compose up --build
```

---

## API Endpoints

### User Management
| HTTP Method | Endpoint       | Description              |
|-------------|----------------|--------------------------|
| `POST`      | `/users`       | Create a new user        |

### Group Management
| HTTP Method | Endpoint       | Description            |
|-------------|----------------|------------------------|
| `POST`      | `/groups`      | Create a new group     |
| `GET`       | `/groups`      | Retrieve all groups    |
| `PUT`       | `/groups`      | Update a group         |
| `GET`       | `/groups/{id}` | Retrieve group details |
| `DELETE`    | `/groups/{id}` | Delete a group         |

### Member Management
| HTTP Method | Endpoint       | Description             |
|-------------|----------------|-------------------------|
| `POST`      | `/member`      | Create a new member     |
| `GET`       | `/member`      | Retrieve all member     |
| `PUT`       | `/member`      | Update a member         |
| `GET`       | `/member/{id}` | Retrieve member details |
| `DELETE`    | `/member/{id}` | Delete a member         |

### Billing
| HTTP Method | Endpoint      | Description           |
|-------------|---------------|-----------------------|
| `POST`      | `/bill`       | Create a new bill     |
| `GET`       | `/bill`       | Retrieve all bill     |
| `PUT`       | `/bill`       | Update a bill         |
| `GET`       | `/bill/{id}`  | Retrieve bill details |
| `DELETE`    | `/bill/{id}`  | Delete a bill         |

### Authentication
| HTTP Method | Endpoint             | Description                    |
|-------------|----------------------|--------------------------------|
| `POST`      | `/authenticate`      | Authenticate and receive JWT   |

### BillCalculation
| HTTP Method | Endpoint               | Description        |
|-------------|------------------------|--------------------|
| `GET`       | `/calculate/{groupId}` | Calculate Expenses |

---

## Configuration
- **Database:** Ensure PostgreSQL is running locally or configure the `application.properties` for a remote instance.
- **Environment Variables:**
    - `DB_USERNAME`: Database username.
    - `DB_PASSWORD`: Database password.
    - `JWT_SECRET`: Secret key for JWT generation.

---

## 📞 Contact Us
For any questions or issues, feel free to reach out:

| Name              | Contact Info                                                           |
|-------------------|------------------------------------------------------------------------|
| Mojtaba Zamandi   | mojtabazamandi.mz@gmail.com                                            |
| GitHub            | [Mojtaba98mz](https://github.com/Mojtaba98mz)                          |
| GitHub Issues     | [Open an Issue](https://github.com/Mojtaba98mz/BillManagement/issues)  |
| LinkedIn          | [Mojtaba Zamandi](https://linkedin.com/in/mojtaba-zamandi)             |

---

## License
## License

This project is licensed under the MIT License - see the full text below.

MIT License

Copyright (c) [2025] [Mojtaba Zamandi]

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