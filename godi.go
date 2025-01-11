package godi

import (
	"fmt"
	"reflect"
)

// Container holds all the registered dependencies and their container names
type Container struct {
	instances map[string]interface{}
	name      string
}

// NewContainer creates a new DI container with an optional name
func NewContainer(name string) *Container {
	return &Container{
		instances: make(map[string]interface{}),
		name:      name,
	}
}

// Get returns a specific dependency from the container
func (c *Container) Get(name string) (interface{}, bool) {
	instance, exists := c.instances[name]
	return instance, exists
}

// Register binds a specific implementation to a name in the container
func (c *Container) Register(name string, instance interface{}) {
	if _, exists := c.instances[name]; exists {
		fmt.Printf("Warning: Dependency %s is already registered. Overwriting.\n", name)
	}
	c.instances[name] = instance
}

// Resolve injects the dependencies into the target struct
func (c *Container) Resolve(target interface{}) error {
	val := reflect.ValueOf(target)
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("target must be a pointer to a struct")
	}

	// Iterate over the struct fields
	val = val.Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		injectTag := field.Tag.Get("inject")
		if injectTag == "" {
			continue
		}

		// Find the dependency by the tag name
		dependency, exists := c.instances[injectTag]
		if !exists {
			return fmt.Errorf("dependency %s not found in container %s", injectTag, c.name)
		}

		// Set the dependency
		fieldValue := val.Field(i)
		if fieldValue.CanSet() {
			fieldValue.Set(reflect.ValueOf(dependency))
		} else {
			return fmt.Errorf("cannot set dependency %s", injectTag)
		}
	}

	return nil
}

// ResolveAll resolves multiple dependencies and panics if any dependency fails
func (c *Container) ResolveAll(targets ...interface{}) {
	for _, target := range targets {
		if err := c.Resolve(target); err != nil {
			panic(fmt.Sprintf("failed to resolve dependency: %v", err))
		}
	}
}

// ListDependencies returns a string listing all registered dependencies in the container
func (c *Container) ListDependencies() []string {
	if len(c.instances) == 0 {
		return []string{"No dependencies registered."}
	}

	var dependencies []string
	for name, instance := range c.instances {
		dependencies = append(dependencies, fmt.Sprintf("%s: %s", name, reflect.TypeOf(instance)))
	}

	return dependencies
}
