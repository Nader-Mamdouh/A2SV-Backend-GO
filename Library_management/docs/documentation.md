# Library Management System

## Overview
A simple console-based system written in Go to manage books and members in a library.

## Features
- Add and remove books
- Borrow and return books
- List available and borrowed books

## Structure
- `main.go`: Entry point
- `models/`: Book and Member structs
- `services/`: Core business logic (implements `LibraryManager` interface)
- `controllers/`: Console UI logic

## Usage
Run the program and use the console menu to manage the library.

## Future Enhancements
- Persist data to file or database
- Add user authentication
- Improve console UI experience