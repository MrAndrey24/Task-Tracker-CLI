# Todo CLI Application

This is a simple command-line interface (CLI) application for managing a list of todos. The application allows you to add, delete, edit, and list todos, as well as mark them as in progress or done.

## Features

- Add a new todo
- Delete an existing todo
- Edit an existing todo
- Mark a todo as in progress
- Mark a todo as done
- List all todos
- List todos by status (not started, in progress, done)

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/MrAndrey24/Task-Tracker-CLI.git
    ```

2. Build the application:
    ```sh
    go build -o todo-cli
    ```

## Usage

Run the application with the desired command-line flags:

- Add a new todo:
    ```sh
    go run . -add "New Todo Title"
    ```

- Delete a todo by index:
    ```sh
    go run . -del 1
    ```

- Edit a todo by index:
    ```sh
    go run . -update 1:"Updated Todo Title"
    ```

- Mark a todo as in progress by index:
    ```sh
    go run . -progress 1
    ```

- Mark a todo as done by index:
    ```sh
    go run . -done 1
    ```

- List all todos:
    ```sh
    go run . -list
    ```

- List todos that are done:
    ```sh
    go run . -list-done
    ```

- List todos that are in progress:
    ```sh
    go run . -list-progress
    ```

- List todos that are not started:
    ```sh
    go run . -list-not-started
    ```

## Flags

- `-add`: Add a new todo. Specify the title of the todo.
- `-del`: Delete a todo by index.
- `-update`: Edit a todo by index. Specify the index and the new title in the format `index:new_title`.
- `-progress`: Mark a todo as in progress by index.
- `-done`: Mark a todo as done by index.
- `-list`: List all todos.
- `-list-done`: List all todos that are done.
- `-list-progress`: List all todos that are in progress.
- `-list-not-started`: List all todos that are not started.

## Example

```sh
go run . -add "Learn Go"
go run . -list
go run . -update 0:"Learn Go in depth"
go run . -progress 0
go run . -done 0
go run . -list-done
go run . -list-progress
go run . -list-not-started
go run . -del 0