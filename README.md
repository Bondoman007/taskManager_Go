# 📝 Task Manager Microservice (Go) - Alle Backend Assignment

This is a simple Task Management backend service written in Go as part of the Alle Backend Assignment. The service demonstrates microservice principles, clean RESTful API design, and supports pagination, filtering, and validation.

---

## 📦 Features

- CRUD operations for tasks (`Create`, `Read`, `Update`, `Delete`)
- Filter tasks by status (`Pending`, `Completed`)
- Pagination support on task listing
- Input validation for task status
- Follows single-responsibility & microservices principles
- Clean, modular Go project structure

---

## 🧠 Problem Breakdown & Design Decisions

### ✅ Requirements:
- RESTful task service
- Pagination: `/tasks?page=1&limit=5`
- Filtering: `/tasks?status=Completed`
- In-memory storage (no DB needed)
- Clean code structure

### ✅ Design Choices:
- **Go**: Chosen for its simplicity, performance, and microservice readiness.
- **In-Memory Map**: Simulates a database for simplicity.
- **Packages Structure**:
  - `models`: Task struct definitions
  - `handlers`: Handle HTTP requests
  - `services`: Business logic
  - `routes`: Route declarations
  - `db`: In-memory mock DB
- **Validation**: Only `"Pending"` and `"Completed"` statuses allowed.

---

## 📁 Project Structure
```bash
taskManager_Go/
├── db/
├── handlers/
├── models/
├── routes/
├── services/
├── main.go
└── README.md
```
## 🚀 Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/Bondoman007/taskManager_Go.git
cd taskManager_Go
```
### 2. Initialize Go Modules
```bash
go mod tidy
```
### 3. Run the Server
```bash
go run main.go
```


### 📡 API Documentation
| Method | Endpoint                 | Description                                               |
| ------ | ------------------------ | --------------------------------------------------------- |
| GET    | `/`                      | Welcome route                                             |
| POST   | `/tasks`                 | Create a new task                                         |
| GET    | `/tasks`                 | Get all tasks (supports pagination & filtering by status) |
| GET    | `/tasks/{id}`            | Get a task by ID                                          |
| PUT    | `/tasks/{id}`            | Update a task by ID                                       |
| DELETE | `/tasks/{id}`            | Delete a task by ID                                       |
| GET    | `/tasks/status/{status}` | Get tasks by status only (`Pending` or `Completed`)       |

### 🧱 Microservices Concepts Demonstrated
🔹 Single Responsibility
-Each Go package handles one responsibility (routing, service logic, etc.).

🔹 Scalable Architecture
-Stateless design (easy to scale horizontally).

-Real database can replace in-memory map easily.

-REST APIs designed for easy consumption and expansion.

### ✨ Future Improvements
Add a real database (e.g., PostgreSQL)

Add user authentication

Deploy with Docker + Kubernetes

Add unit & integration tests

### 👨‍💻 Author
Made with ❤️ by Kanishk Shrivastava 
