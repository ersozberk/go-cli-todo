# 🚀 Go REST CLI: Todo Fetcher

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)
![Platform](https://img.shields.io/badge/Platform-Linux%20%7C%20macOS%20%7C%20Windows-lightgrey)
![License](https://img.shields.io/badge/License-MIT-green)

A fast, dependency-free command-line interface tool built with Go. This project demonstrates core language mechanics, network requests, and data modeling by fetching data from a REST API and formatting it in the terminal.

## ✨ Features

* **Zero Dependencies:** Built entirely using Go's standard library (`net/http`, `encoding/json`, `flag`).
* **Dynamic Parameter Handling:** Generates dynamic API requests via command-line flags.
* **Type-Safe Data Processing:** Securely unmarshals raw JSON responses into strict Go `struct` types.
* **Standard Go Project Layout:** Adheres to the community-standard `cmd/` directory structure for maintainability.

## 📂 Project Structure

```text
.
├── cmd/
│   └── todo-getir/
│       └── main.go       # Application entry point and core logic
├── go.mod                # Go module and dependency declarations
├── .gitignore            # Files ignored by Git
└── README.md             # Project documentation


🛠 Installation and Usage
Prerequisites
Go installed on your machine.

Running from Source
To clone the repository and run the code directly without building:

```Bash
git clone [https://github.com/ersozberk/go-cli-todo.git](https://github.com/ersozberk/go-cli-todo.git)
cd go-cli-todo
go run cmd/todo-getir/main.go --id=5
Building for Production
To compile the Go code into a standalone, executable binary:

Bash
# Compile the project
go build -o todo-getir cmd/todo-getir/main.go

# Execute the binary
./todo-getir --id=12
(Optional) Move the executable to your system path to run it from anywhere (Linux/macOS):

```Bash
sudo mv todo-getir /usr/local/bin/
todo-getir --id=42
🧠 Learning Outcomes
This project serves as a foundational exercise in mastering Go, focusing on:

Module initialization and package management (go mod).

Memory addresses (Pointers) and Struct manipulation.

Resource management and preventing memory leaks using the defer keyword.

Parsing terminal arguments via the flag package.
