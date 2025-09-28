## CRUD Operation in golang

- this project provide api for todo list.
- you can list, add, delete and do mark as complete of todos.
- we use mongodb for database.

## Setup

- clone repository

```bash
https://github.com/amitkys/goCRUD.git
```

- install dependencies

```bash
go mod tidy
```

- setup environment variables

by looking at `.env.example` file, create a `.env` file and add your environment variables.

## Endpoints

### `GET /api/todos`

- Get all todos

### `POST /api/todos`

- Create a new todo
  data - Request Body: `{"body":"go to gym"}`

### `PUT /api/todos/{id}`

- Update an existing todo of it's completion status

### `DELETE /api/todos/{id}`

- Delete a todo
