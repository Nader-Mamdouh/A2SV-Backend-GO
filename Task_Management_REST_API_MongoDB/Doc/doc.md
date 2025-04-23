# Task Management API Documentation

## Overview

This document provides comprehensive information about the Task Management API, enhanced with persistent data storage using **MongoDB** and the official **MongoDB Go Driver**. This enhancement replaces the previous in-memory storage, ensuring data persistence across API restarts.

---

## 📌 Objective

To integrate MongoDB as the persistent data layer for managing tasks in the Task Management API, allowing full CRUD operations with robust data validation and error handling.

---

## 🛠️ Technology Stack

- **Go (Golang)**: Backend development
- **MongoDB**: NoSQL database for persistent storage
- **MongoDB Go Driver**: Official driver to interact with MongoDB
- **Gin**: Web framework for routing and HTTP handling

---

## 🗂️ Folder Structure

task_manager/ ├── main.go # Entry point of the application ├── controllers/ │ └── task_controller.go # Handles HTTP requests ├── models/ │ └── task.go # Defines Task struct ├── data/ │ └── task_service.go # MongoDB logic using Go Driver ├── router/ │ └── router.go # Route setup using Gin ├── docs/ │ └── api_documentation.md # API and MongoDB integration documentation └── go.mod # Go module and