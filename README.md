# GoDI: Lightweight Dependency Injection for Go

`godi` is a simple and lightweight dependency injection (DI) library for Go. It allows developers to manage dependencies in a structured and centralized way, making your code more modular and testable.

## Features

- **Centralized Dependency Management**: Register and manage dependencies in a container.
- **Dependency Injection**: Automatically inject dependencies into structs using tags.
- **Multi-Dependency Resolution**: Resolve multiple dependencies in a single call.
- **Debugging**: List all registered dependencies for introspection and debugging.

## Installation

To install `godi`, use `go get`:

```bash
go get github.com/zombocoder/godi
```

## Usage

### 1. Create a Container

Create a DI container to manage your dependencies:

```go
import "github.com/zombocoder/godi"

container := godi.NewContainer("MyContainer")
```

### 2. Register Dependencies

Register your dependencies in the container:

```go
type MyService struct{}

container.Register("MyService", &MyService{})
```

### 3. Inject Dependencies

Use struct tags to specify which dependency to inject:

```go
type MyStruct struct {
    Service *MyService `inject:"MyService"`
}

instance := &MyStruct{}
err := container.Resolve(instance)
if err != nil {
    log.Fatalf("failed to resolve dependencies: %v", err)
}

// Now instance.Service is automatically populated
```

### 4. Resolve Multiple Dependencies

You can resolve multiple dependencies at once:

```go
anotherInstance := &AnotherStruct{}
container.ResolveAll(instance, anotherInstance)
```

### 5. List Registered Dependencies

Debug or introspect your container by listing all registered dependencies:

```go
dependencies := container.ListDependencies()
for _, dep := range dependencies {
    fmt.Println(dep)
}
```

## Advanced Usage

### Overwriting Dependencies

You can overwrite an existing dependency by registering it again:

```go
type AnotherService struct{}

container.Register("MyService", &AnotherService{})
```

### Handling Missing Dependencies

If a dependency is not registered, the `Resolve` method will return an error:

```go
type MissingStruct struct {
    MissingDep *MissingService `inject:"MissingService"`
}

err := container.Resolve(&MissingStruct{})
if err != nil {
    fmt.Println("Error:", err)
}
```

### Error Handling for Unsettable Fields

Ensure all fields to be injected are exported; otherwise, the `Resolve` method will return an error.

## Contributing

Contributions are welcome! Feel free to fork this repository, open an issue, or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
