package main

type Service struct {
	Repository Repository `inject:"InMemoryRepository"`
	Storage    Storage    `inject:"InMemoryStorage"`
	Logger     Logger     `inject:"LogrusLogger"`
}

// NewService initializes a new Service
func NewService() *Service {
	return &Service{}
}

// GetByID fetches an entity by its ID
func (s *Service) GetByID(id int) string {
	s.Logger.Info("Downloading file", map[string]interface{}{"url": "http://example.com"})
	s.Storage.Download("http://example.com")
	s.Logger.Info("Get repository by ID", map[string]interface{}{"id": id})
	return s.Repository.GetByID(id)
}
