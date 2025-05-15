# API_CRUD

A simple CRUD API for user management built using Golang and the [go-chi](https://github.com/go-chi/chi) framework.

## Features

- Create, Read, Update, and Delete (CRUD) operations for users
- Lightweight and efficient API built with Go
- Uses `go-chi` for routing

## Getting Started

### Installation

Clone the repository:

```bash
git clone https://github.com/joaopedroldavid-del/API_CRUD.git
cd API_CRUD
```

### Running the API
To start the API server, run:

```sh
go run main.go
```

### API Endpoints
### Method / Endpoint / Description
- GET	/users	Retrieve all users
- GET	/users/{id}	Retrieve a user by ID
- POST	/users	Create a new user
- PUT	/users/{id}	Update a user
- DELETE	/users/{id}	Delete a user