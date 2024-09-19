# Prerequisite

1. Make sure you have the following installed outside the current project directory and available in your `GOPATH`
    - golang
    - [air](https://github.com/air-verse/air) for hot reloading
    - `.mysql.env` is sample database
    - `docker` or `docker-compose`

# Installation

1. Clone this repo
2. Run `go get`
3. 
```shell
    $ make up
```

# Running

1. Type `air` in the command line

# Environment Variables

2. Endpoints  

GET /api/v1/healthcheck - HealthCheck

POST /api/v1/users - Create User  
POST /api/v1/users/login - Login user  
PUT /api/v1/users/:id - Update user  
PATCH /api/v1/users/:id - Update User with partial  

GET /api/v1/memos - Get all memos  
POST /api/v1/memos - Create memo  
GET /api/v1/memos/:id - Get memo  
PUT /api/v1/memos/:id - Update memo  
PATCH /api/v1/memos/:id - Update memo partial  
DELETE /api/v1/memos/:id - Delete memo