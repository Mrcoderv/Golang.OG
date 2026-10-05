// dependency inversion 

type UserRepository interface {
    Save(name string) error
}

type UserService struct {
    repo UserRepository
}

func (s UserService) CreateUser(name string) error {
    return s.repo.Save(name)
}



//using the postgres repository
type PostgresRepository struct{}

func (p PostgresRepository) Save(name string) error {
    // PostgreSQL logic
    return nil
}

// using the post grace repo 
 type MySQLRepository struct{}

func (m MySQLRepository) Save(name string) error {
	// MySQL logic
	return nil
}