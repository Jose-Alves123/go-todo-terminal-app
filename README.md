# GoDoIt

This terminal application aims to provide aid for developers to schedule their tasks.

To run the application, in the project root directory, run:

```bash
go run ./cmd/app/main.go
```

## Features

- ✅ Create ToDo
- ✅ Edit ToDo
- ✅ Delete ToDo
- ✅ List ToDo
- ✅ Allow delete local database
- ❌ Create new comments on ToDo
- ❌ Allow user to add new states of ToDo (besides the base states To Do, Doing, Done)
- ❌ Create new tags (ex: work, personal, adventure, misc, gaming, programming)
- ❌ Create new groups (ex: work, personal, adventure, misc, gaming, programming)
- ❌ Filter ToDo based on state, groups and / or tags
- ❌ Allow ToDo to have up to one group
- ❌ Allow ToDo to have n tags
- ❌ In application show calendar of states changes of todos
- ❌ Show graphics of states of ToDos over time
- ❌ Create command to allow info to be shown on terminal, every time is started, about the number of incomplete ToDos
- ❌ Add Locatization for English and Portuguese, allowing user to use language of choice

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

The TUI is using [bubbletea](https://github.com/charmbracelet/bubbletea), with lipgloss and bubbles.
