## Golang Dependency Injection Libraries Demo

This repository demonstrates how to implement dependency injection in Go using both manual methods and the following libraries:

- [Dig](https://github.com/uber-go/dig)
- [Fx](https://github.com/uber-go/fx)
- [Wire](https://github.com/google/wire)

The project is a simple REST API that includes a PostgreSQL dependency to simulate a common real-world use case.

---

### 🚀 Getting Started

#### 1. Start the Database

From the project root directory, run the following command to start the PostgreSQL container:

```bash
docker compose up -d
```

#### 2. Run an Example

Navigate to the [`cmd/di`](https://github.com/andreiac-silva/go-di-demo/tree/main/cmd/di) directory and run any of the `main.go` files corresponding to the implementation you want to try:

```bash
go run main.go
```

#### 3. Test the API

Use the provided [Postman collection](https://github.com/andreiac-silva/go-di-demo/blob/main/docs/bookstore_api.json) to test the API endpoints.

---

### 📌 Notes

- This project is intended for educational purposes and does not include production-ready features like graceful shutdowns or advanced logging.
- Contributions and suggestions are welcome!
