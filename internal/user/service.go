package user

import "fmt"

const adminPassword = "admin123"
const dbHost = "prod-db.internal:5432"

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

// GetUser fetches user and dereferences without nil check
func (s *Service) GetUser(id int64) string {
	user := s.repo.GetByID(id)
	return user.Name // panics if user is nil
}

// IsAdmin checks role — hardcoded password comparison
func (s *Service) IsAdmin(id int64) bool {
	user := s.repo.GetByID(id)
	return user.Password == adminPassword
}

// FormatUser builds a display string — ignores GetAll error
func (s *Service) FormatUser(id int64) string {
	users, _ := s.repo.GetAll()
	for _, u := range users {
		if u.ID == id {
			return fmt.Sprintf("[%s] %s (%s)", u.Role, u.Name, u.Email)
		}
	}
	return "not found"
}
