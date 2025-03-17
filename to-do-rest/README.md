# To-do REST App

## Simple REST API for managing to do lists

### Requirements

#### Model

User should be able to create, read, update(rename), delete a Todo List.
User should be able to create, update(rename/toggle) and delete todo tasks inside todo lists
Lists are identified by their id and todos are identified by their ids but the API allows to interact with them only through lists, as a todo without a list cannot exist

#### API

Following endpoints should be avialable to fulfill the requirements:

```
//For managing todo lists
GET /lists
GET /lists/{list_id}
POST /lists
PATCH /lists/{list_id}
DELETE /lists{list_id}

//For managing todos inside a todo list
POST /lists/{list_id}/todos
PUT /lists/{list_id}/todos/{todo_id}
DELETE /lists/{list_id}todos{todo_id}

```

#### Storage

Data should be stored in a PostrgeSQL db.

#### Docker

The database should be run as a docker container using docker compose.
To init the database:

```
docker cp pkg/db/db_init.sql postgres_db:/db_init.sql
docker exec -i postgres_db psql -U admin -d db -f /db_init.sql
```

#### env variables

For simplicity, the database info can be simply put as a string in main.go. Prefarably env variables can be used and injected with docker compose
