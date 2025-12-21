package user

type Service interface {
	ListUsers() []User
}

type Services struct {
	repo Repository
}

// ListUsers implements [Service].
func (s Services) ListUsers() []User {
	return s.repo.GetAll()
}

func NewService(repo Repository) Services {
	return Services{repo: repo}
}
