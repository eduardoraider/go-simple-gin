# Simple Go Gin API

This project utilizes Gin, a popular HTTP web framework for Go, to create a basic API server.

## Setup

To install Gin, run the following command:

```
go get -u github.com/gin-gonic/gin
```

To run the server, execute:

```
go run main.go
```

## API Endpoints

- **GET /todos**: Retrieve all todo items.
- **GET /todos/:id**: Retrieve a todo item by its ID.
- **POST /todos/add**: Add a new todo item.

## Usage

### Using CLI (curl)

Get all todos:
```
curl http://localhost:8080/todos
```

Get a todo by ID:
```
curl http://localhost:8080/todos/123
```

Add a new todo:
```
curl -X POST -H "Content-Type: application/json" -d '{"title": "New Todo", "done": false}' http://localhost:8080/todos/add
```

### Using Postman

Get all todos:
- Method: GET
- URL: http://localhost:8080/todos


Get a todo by ID:
- Method: GET
- URL: http://localhost:8080/todos/123
- Replace 123 with the ID of the todo you want to retrieve.


Add a new todo:
- Method: POST
- URL: http://localhost:8080/todos/add
- Headers:
    - Key: Content-Type
    - Value: application/json
- Body (raw JSON):
  ```
  {
      "title": "New Todo",
      "done": false
  }
  ```


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

by **Eduardo Raider** - Software Engineer