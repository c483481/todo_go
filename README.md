# Simple ToDo List Backend With GO

## introduction

This is a simple ToDo List backend application built with Go. It provides a REST API for managing ToDo items, allowing users to create, read, update, and delete tasks. This project demonstrates the basics of building a REST API with Go, using clean architecture principles.

## Feature
- Create, Read, Update, and Delete (CRUD) ToDo items
- Simple and clean code structure
- Uses the Fiber framework for handling HTTP requests

## Installation

1. Clone the repository
    ```shell
    git clone https://github.com/c483481/todo_go
    cd todo_go
    ```
2. Install dependencies
    ```shell
    go mod tidy
    ```
3. Setup environment variable
   - Create a .env file in the root of the project and add your configuration settings.
   - see the example configuration file from the .enx.example file
4. Run the application
    ```shell
    make run
    ```