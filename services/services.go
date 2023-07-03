package services

import repositorys "main/repository"

//go:generate mockgen -destination=../mocks/service/mockTodoservice.go -package=services main/services TodoService
type UserService struct {
	userRepository repositorys.UserRepository
}
type TodoService interface {
	AddUser(name string) error
	GetUserByID(id int) (string, error)
}

func NewUserService(userRepository repositorys.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) AddUser(name string) error {
	// Gerekli iş mantığını uygula
	return s.userRepository.AddUser(name)
}

func (s *UserService) GetUserByID(id int) (string, error) {
	// Gerekli iş mantığını uygula
	return s.userRepository.GetUserByID(id)
}
