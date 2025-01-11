// Package godi provides a simple and lightweight dependency injection container
// for Go applications. It allows developers to register dependencies, resolve
// them into structs, and manage their lifecycle in a centralized and organized manner.
//
// Features:
//   - Register and manage named dependencies.
//   - Inject dependencies into struct fields using struct tags.
//   - Resolve multiple dependencies at once.
//   - List all registered dependencies for debugging and introspection.
//
// Usage:
//
// 1. Creating a Container
//
//	// Create a new DI container
//	container := godi.NewContainer("MyContainer")
//
// 2. Registering Dependencies
//
//	// Register a dependency with a name
//	container.Register("MyService", &MyService{})
//
// 3. Resolving Dependencies
//
//	type MyStruct struct {
//	    Service *MyService `inject:"MyService"`
//	}
//
//	instance := &MyStruct{}
//	err := container.Resolve(instance)
//	if err != nil {
//	    log.Fatalf("failed to resolve dependencies: %v", err)
//	}
//
// 4. Resolving Multiple Dependencies
//
//	anotherInstance := &AnotherStruct{}
//	container.ResolveAll(instance, anotherInstance)
//
// 5. Listing Dependencies
//
//	dependencies := container.ListDependencies()
//	fmt.Println("Registered dependencies:", dependencies)
package godi
