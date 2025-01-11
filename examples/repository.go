package main

// Repository defines the interface for a repository
type Repository interface {
	GetByID(id int) string
}

// InMemoryRepository is a simple implementation of the Repository interface
type InMemoryRepository struct {
	Logger   Logger `inject:"LogrusLogger"`
	entities map[int]string
}

// NewInMemoryRepository creates a new InMemoryRepository
func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		entities: map[int]string{
			1: "Entity One",
			2: "Entity Two",
		},
	}
}

// GetByID fetches an entity by ID
func (r *InMemoryRepository) GetByID(id int) string {
	// r.Logger.Info("Fetching entity", map[string]interface{}{"id": id})
	return r.entities[id]
}

type SomeOtherRepository struct {
	Logger   Logger `inject:"LogrusLogger"`
	entities map[int]string
}

func NewSomeOtherRepository() *SomeOtherRepository {
	return &SomeOtherRepository{
		entities: map[int]string{
			1: "Some other Entity One",
			2: "Some other Entity Two",
		},
	}
}

func (r *SomeOtherRepository) GetByID(id int) string {
	// r.Logger.Info("Fetching some other entity", map[string]interface{}{"id": id})
	return r.entities[id]
}
