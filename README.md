# Task Manager Command Line

## Description

This is a simple command line task manager that allows you to add, remove, and list tasks. It is written in Go and uses a SQLite database to store tasks. Its uses cobras CLI library to create the command line interface with tview as the command line interface.

## Installation

To install the task manager, you can clone the repository and run the following command:

```bash
go install
```

## Usage

To use the task manager, you can run the following command:

```bash
task-manager
```

This will display the help menu with all the available commands.

## Commands

The task manager has the following commands:

- `init`: Initialize the task manager
- `add`: Add a new task
- `remove`: Remove a task
- `list`: List all tasks
- `complete`: Mark a task as complete
- `incomplete`: Mark a task as incomplete

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Author

- Erwin Komaruloh
