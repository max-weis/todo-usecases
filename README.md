# Todo Use Cases Library

This project provides a collection of use cases for a simple Todo application, designed to be used as a benchmark or a test suite for various software libraries and frameworks. 
It includes fundamental operations for managing todos, such as creation, listing, updating, deleting, toggling completion status, and searching. 
The use cases are implemented in a way that allows for easy integration with different persistence layers and UI frameworks, making it an excellent tool for:

- **Testing** new libraries or frameworks without the overhead of writing complex business logic.
- **Prototyping** new ideas or features in the context of a familiar todo application.

## Use Cases

The following use cases are implemented in this library:

- **Create Todo**: Add a new todo item to the system.
- **List Todos**: Retrieve todos with pagination support.
- **Toggle Todo**: Change the completion status of a todo.
- **Delete Todo**: Remove a todo from the system.
- **Search Todo**: Search for todos based on a given query.

Each use case is designed to be independent, allowing for modular testing and integration with different backends:

### Create Todo
- Functionality: Create a new todo item with a title and due date.
- Validation: Checks for non-empty titles and future due dates.

### List Todos
- Functionality: Retrieves a paginated list of todos.
- Features: Supports filtering by completion status and pagination.

### Toggle Todo
- Functionality: Toggle the completion status of a todo item.

### Delete Todo
- Functionality: Remove a todo from the system by its ID.
- Validation: Ensures the ID is not empty.

### Search Todo
- Functionality: Search for todos based on text input.
- Flexibility: Can be adapted to search by title, description, or tags (when implemented).

## How to Use

### Installation

To get started with the library, clone this repository or add it as a submodule to your project:

```sh
git clone https://github.com/max-weis/todo-usecases
```

Or if you are using Go modules:

```sh
go get https://github.com/max-weis/todo-usecases
```

### Implementing in Your Project

1. **Import the Package**:
   ```go
   import "https://github.com/max-weis/todo-usecases/create"
   ```

2. **Create a Repository or Service Layer**:
   - Implement the interfaces or functions required by each use case (like `SaveTodo`, `GetTodosWithPagination`, etc.)

3. **Instantiate Use Cases**:
   - Use the `New*` functions provided in the library to create use case instances, passing in your repository/service implementations.

   ```go
   saveTodoFunc := func(ctx context.Context, todo todo.Todo) error {
       // Your implementation to save a todo
   }
   createTodo := todo.NewCreateTodoUseCase(saveTodoFunc)
   ```

4. **Use the Use Cases**:
   - Call the use case functions with the necessary parameters.

   ```go
   newTodo, err := createTodo(context.Background(), "Write Readme", time.Now().Add(24*time.Hour))
   ```

## Testing

- Tests are written for each use case to ensure correct behavior.
- You can run the tests by executing:

```sh
go test ./...
```
