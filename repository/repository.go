package repositorys

//go:generate mockgen -destination=../mocks/repository/mockTodoRepository.go -package=repository main/repository UserRepository
type UserRepositoryImpl struct {
	// Veritabanı bağlantısı veya başka bir depolama alanı
}

type UserRepository interface {
	AddUser(name string) error
	GetUserByID(id int) (string, error)
}

func (r *UserRepositoryImpl) AddUser(name string) error {
	// Kullanıcıyı ekleme işlemi
	return nil
}

func (r *UserRepositoryImpl) GetUserByID(id int) (string, error) {
	// ID'ye göre kullanıcıyı getirme işlemi
	return "", nil
}
