package godi_test

import (
	"testing"

	"github.com/zombocoder/godi"
)

type MockDependency struct{}

type MockService struct {
	Dependency *MockDependency `inject:"MockDependency"`
}

type SomeService struct {
	dependency *MockDependency `inject:"MockDependency"` // Unexported field
}

type SomeNotInjectedService struct {
	dependency *MockDependency `inject:""` // No inject tag specified
}

func TestGet(t *testing.T) {
	// Test the Get method to retrieve registered dependencies
	container := godi.NewContainer("TestContainer")
	mockDependency := &MockDependency{}
	container.Register("MockDependency", mockDependency)

	// Verify that a registered dependency can be retrieved
	instance, exists := container.Get("MockDependency")
	if !exists {
		t.Errorf("expected dependency 'MockDependency' to exist")
	}
	if instance != mockDependency {
		t.Errorf("expected returned dependency to match the registered instance")
	}

	// Verify that retrieving a non-existent dependency returns false
	_, exists = container.Get("NonExistentDependency")
	if exists {
		t.Errorf("did not expect 'NonExistentDependency' to exist")
	}
}

func TestResolve(t *testing.T) {
	// Test the Resolve method to inject dependencies into struct fields
	container := godi.NewContainer("TestContainer")
	mockDependency := &MockDependency{}
	container.Register("MockDependency", mockDependency)

	// Verify that dependencies are injected correctly
	mockService := &MockService{}
	err := container.Resolve(mockService)
	if err != nil {
		t.Errorf("expected no error when resolving dependencies, got %s", err)
	}

	if mockService.Dependency != mockDependency {
		t.Errorf("expected resolved dependency to match the registered instance")
	}

	// Verify that resolving a non-pointer target returns an error
	err = container.Resolve(mockDependency)
	if err == nil {
		t.Errorf("expected error when resolving a non-pointer target, got nil")
	}

	// Verify behavior when the inject tag is empty
	someNotInjectedService := &SomeNotInjectedService{}
	err = container.Resolve(someNotInjectedService)
	if err != nil {
		t.Errorf("expected no error when resolving dependencies with empty inject tag, got %s", err)
	}

	// Verify error handling for non-struct pointers
	someFunc := func(t *testing.T) {}
	err = container.Resolve(&someFunc)
	if err == nil {
		t.Errorf("expected error when resolving a non-struct pointer, got nil")
	}
}

func TestRegisterOverwrite(t *testing.T) {
	// Test that Register allows overwriting existing dependencies
	container := godi.NewContainer("TestContainer")
	mockDependency := &MockDependency{}
	anotherDependency := &MockDependency{}

	// Register initial dependency
	container.Register("MockDependency", mockDependency)

	// Overwrite the dependency
	container.Register("MockDependency", anotherDependency)

	// Verify that the dependency is overwritten
	instance, _ := container.Get("MockDependency")
	if instance != anotherDependency {
		t.Errorf("expected dependency to be overwritten with the new instance")
	}
}

func TestResolveUnsettableField(t *testing.T) {
	// Test resolving dependencies for fields that cannot be set
	container := godi.NewContainer("TestContainer")
	mockDependency := &MockDependency{}
	container.Register("MockDependency", mockDependency)

	unsettableService := &SomeService{}
	err := container.Resolve(unsettableService)
	if err == nil {
		t.Fatalf("expected error when resolving unsettable field, got nil")
	}

	// Verify that the error message is correct
	expectedError := "cannot set dependency MockDependency"
	if err.Error() != expectedError {
		t.Errorf("expected error '%s', got '%s'", expectedError, err.Error())
	}
}

func TestResolveAllWithPanic(t *testing.T) {
	// Test that ResolveAll panics when dependencies are unresolved
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic when resolving unresolved dependencies")
		}
	}()

	container := godi.NewContainer("TestContainer")
	mockService := &MockService{}
	container.ResolveAll(mockService)
}

func TestListDependencies(t *testing.T) {
	// Test that ListDependencies returns the correct list of registered dependencies
	container := godi.NewContainer("TestContainer")
	mockDependency := &MockDependency{}
	container.Register("MockDependency", mockDependency)

	dependencies := container.ListDependencies()
	if len(dependencies) != 1 {
		t.Errorf("expected 1 dependency, got %d", len(dependencies))
	}

	// Verify the dependency format
	expected := "MockDependency: *godi_test.MockDependency"
	if dependencies[0] != expected {
		t.Errorf("expected '%s' in dependencies, got '%s'", expected, dependencies[0])
	}
}

func TestListNoDependencies(t *testing.T) {
	// Test that ListDependencies correctly reports when no dependencies are registered
	container := godi.NewContainer("TestContainer")

	dependencies := container.ListDependencies()
	if len(dependencies) != 1 {
		t.Errorf("expected 1 dependency, got %d", len(dependencies))
	}
	if dependencies[0] != "No dependencies registered." {
		t.Errorf("expected 'No dependencies registered.' in dependencies, got '%s'", dependencies[0])
	}
}
