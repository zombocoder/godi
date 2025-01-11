package main

type Storage interface {
	Download(url string) error
}

// InMemoryStorage is a simple implementation of the Storage interface
type InMemoryStorage struct{}

// NewInMemoryStorage creates a new InMemoryStorage
func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{}
}

// Download fetches a file from a URL
func (s *InMemoryStorage) Download(url string) error {
	println("Downloading file from", url)
	return nil
}
