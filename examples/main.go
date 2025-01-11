package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/zombocoder/godi"
)

func main() {
	// Logger

	// Create the DI container
	container := godi.NewContainer("main")

	// Register dependencies
	container.Register("SomeOtherRepository", NewSomeOtherRepository())
	container.Register("InMemoryRepository", NewInMemoryRepository())
	container.Register("InMemoryStorage", NewInMemoryStorage())
	container.Register("LogrusLogger", NewLogrusLogger(logrus.DebugLevel, &logrus.JSONFormatter{}))

	// Create the service and resolve dependencies
	service := NewService()
	container.ResolveAll(service, NewInMemoryRepository())
	// container.ValidateDependencies()
	dependencies := container.ListDependencies()
	fmt.Println(dependencies)
	// Use the service
	fmt.Println(service.GetByID(1))
}
