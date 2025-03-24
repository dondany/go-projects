# To-do cli App

## Simple to-do command line tool for task management using GORM

### Requirements

The app should allow basic operations like: to-do item creation, listing, deletion and toggle. For example:

```
$ to-do add <task description>
$ to-do list
$ to-do delete <task id>
$ to-do toggle <task-id>
```

App should use postgresql running on docker as underlying db
To run db container use

```
docker compose up -d
```
