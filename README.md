# Go Todo REST API Example

A RESTful API example for simple todo application with Go

It is a simple tutorial or example for making simple RESTful API with Go using **gorilla/mux** (A nice mux library) and **gorm** (An ORM for Go)

## Installation & Run

### Local Development

```bash
# Clone this project
git clone https://github.com/YourUsername/go-todo-rest-api-example.git
cd go-todo-rest-api-example

# Install dependencies
go mod download

# Run locally
go run main.go

# API Endpoint : http://localhost:8080
```

### Docker

```bash
# Build Docker image
docker build -t todo-api .

# Run container
docker run -p 8080:8080 todo-api

# API Endpoint : http://localhost:8080
```

### Google Cloud Run Deployment

1. Connect your GitHub repository to Cloud Run
2. Set up Cloud Build trigger:
   - Source: GitHub repository
   - Branch: main (or your default branch)
   - Configuration: Dockerfile
   - Location: Repository
3. Push changes to GitHub to trigger automatic deployment
4. Access your API at the provided Cloud Run URL

## Configuration

The application uses environment variables for configuration:

- `PORT`: Server port (defaults to 8080)

Database configuration in [config/config.go](config/config.go):

```go
func GetConfig() *Config {
    return &Config{
        DB: &DBConfig{
            Dialect:  "mysql",
            Username: "guest",
            Password: "Guest0000!",
            Name:     "todoapp",
            Charset:  "utf8",
        },
    }
}
```

## Project Structure

```
├── app
│   ├── app.go
│   ├── handler          // Our API core handlers
│   │   ├── common.go    // Common response functions
│   │   ├── projects.go  // APIs for Project model
│   │   └── tasks.go     // APIs for Task model
│   └── model
│       └── model.go     // Models for our application
├── config
│   └── config.go        // Configuration
├── Dockerfile          // Docker configuration
├── .dockerignore      // Docker ignore file
└── main.go            // Application entry point
```

## API Endpoints

### Projects

#### /projects

- `GET` : Get all projects
- `POST` : Create a new project

#### /projects/:title

- `GET` : Get a project
- `PUT` : Update a project
- `DELETE` : Delete a project

#### /projects/:title/archive

- `PUT` : Archive a project
- `DELETE` : Restore a project

### Tasks

#### /projects/:title/tasks

- `GET` : Get all tasks of a project
- `POST` : Create a new task in a project

#### /projects/:title/tasks/:id

- `GET` : Get a task of a project
- `PUT` : Update a task of a project
- `DELETE` : Delete a task of a project

#### /projects/:title/tasks/:id/complete

- `PUT` : Complete a task of a project
- `DELETE` : Undo a task of a project

## Development Status

- [x] Support basic REST APIs
- [x] Organize the code with packages
- [x] Docker support
- [x] Cloud Run deployment setup
- [ ] Support Authentication with user for securing the APIs
- [ ] Make convenient wrappers for creating API handlers
- [ ] Write tests for all APIs
- [ ] Make docs with GoDoc

!!!
