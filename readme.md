# Simple Todo API

This is a simple Todo API project that allows you to manage a list of todos. The project uses Golang for the backend, and the data is stored in a MySQL database, which can be set up using Docker Compose.

### Prerequisites

Before you start, make sure you have the following installed on your machine:

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Golang](https://golang.org/)

### Setup Database with Docker Compose

1. Create a `docker` folder in your project directory.

2. Create a file named `datamaster.yml` inside the `docker` folder with the following content:

3. Open a terminal, navigate to your project directory, and run the following command to start the MySQL container using Docker Compose:

   ```bash
   docker-compose -f docker/datamaster.yml up -d
   ```

   This command will start the MySQL container named `todo-mysql`, set the root password to "root", create a database named `tododb`, and map port 3306.

### Run the Todo API

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/your-username/todo-api.git
   ```

2. Change into the project directory:

   ```bash
   cd todo-api
   ```

3. Build and run the Golang application:

   ```bash
   go run main.go
   ```

   This will start the Todo API server at `http://localhost:8080`.

### API Endpoints

- **GET /api/tasks**: Get all tasks
- **GET /api/tasks/{id}**: Get a specific task by ID
- **POST /api/tasks**: Create a new task
- **PUT /api/tasks/{id}**: Update a task by ID
- **DELETE /api/tasks/{id}**: Delete a task by ID
- **GET /api/tasks/{id}/status?set=true**: Complated a task by ID

### Example Usage

#### Get all tasks

```bash
curl http://localhost:8080/api/tasks
```

#### Get a task

```bash
curl http://localhost:8080/api/tasks/{id}
```

#### Create a new task

```bash
curl -X POST -H "Content-Type: application/json" -d '{{"title": "Self development","description": "Reading 10 pages of a programming book","due_date": "2024-01-10T00:00:00Z"}}' http://localhost:8080/api/tasks
```

#### Update a task

```bash
curl -X PUT -H "Content-Type: application/json" -d '{{"title": "Self development","description": "Reading 10 pages of a programming book","due_date": "2024-01-10T00:00:00Z"}}' http://localhost:8080/api/tasks/{id}
```

#### Complated a task

```bash
curl -X DELETE http://localhost:8080/api/tasks/{id}
```

#### Update a task

```bash
curl -X GET -H "Content-Type: application/json" http://localhost:8080/api/tasks/{id}/status?set=true
```

### Cleanup

1. To stop and remove the MySQL container, run:

   ```bash
   docker-compose -f docker/datamaster.yml down
   ```

This is a basic setup for a Todo API with Golang and MySQL using Docker Compose. Feel free to customize and extend the project based on your requirements.

### Libraries Used

- [validator/v10](https://pkg.go.dev/github.com/go-playground/validator/v10)
- [go-sql-driver/mysql](https://pkg.go.dev/github.com/go-sql-driver/mysql)
- [google/wire](https://pkg.go.dev/github.com/google/wire)
- [julienschmidt/httprouter](https://pkg.go.dev/github.com/julienschmidt/httprouter)
- [spf13/viper](https://pkg.go.dev/github.com/spf13/viper)
- [stretchr/testify](https://pkg.go.dev/github.com/stretchr/testify)
