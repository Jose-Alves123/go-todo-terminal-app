# GoDoIt

This terminal application aims to provide aid to developers to schedule their tasks.

To run the application, in the project root directory, run:

```bash
go run ./cmd/app/main.go
```

## Features

- ✅ create ToDo
- ❌ Edit ToDo
- ❌ Delete ToDo
- ❌ List ToDo
- ❌ create new comments on ToDo
- ❌ allow user to add new states of ToDo (besides the base states To Do, Doing, Done)
- ❌ create new tags (ex: work, personal, adventure, misc, gaming, programming)
- ❌ create new groups (ex: work, personal, adventure, misc, gaming, programming)
- ❌ filter ToDo based on state, groups and / or tags
- ❌ allow ToDo to have up to one group
- ❌ allow ToDo to have n tags
- ❌ in application show calendar of states changes of todos
- ❌ show graphics of states of ToDos over time
- ❌ create command to allow info to be shown on terminal, every time is started, about the number of incomplete ToDos

## Technologies used

- [GO 1.24.4](https://go.dev)
  - [gorm](https://gorm.io)
  - [gorm/sqlite](https://gorm.io)
  - [testing](https://pkg.go.dev/testing)
  - [godotenv](https://github.com/joho/godotenv)

# Unit Testing

To assert the results are expected, development of application is following TDD practices. To run tests, use command

```shell
go test ./... -coverprofile cover.out
```

# Terminal Interface

In the future a terminal interface will be added to facilitate the use of the program. For that, [bubbletea](https://github.com/charmbracelet/bubbletea) will be used.
