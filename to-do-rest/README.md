# To-do REST App

## Simple REST API for managing to do lists

### Requirements

User should be able to create, read, update(rename), delete a Todo List.
User should be able to create, update(rename/toggle) and delete todo tasks inside todo lists
Lists are identified by their name and todos are identified by the list's name and id

```
Following endpoints should be avialable to fulfill the requirements:

//For managing todo lists
GET /lists
GET /lists/{name}
POST /lists
PATCH /lists/{name}
DELETE /lists{name}

//For managing todos inside a todo list
POST /lists/{name}/todos
PUT /lists/{name}/todos/{id}
DELETE /lists/{name}todos{id}

```
