package user

type Repository interface {
	GetAll() []User
}

type memoryRepo struct {
	users []User
}

func NewMemoryRepo() Repository {
	return &memoryRepo{
		users: []User{
			{ID: 1, Name: "Alice", Role: RoleAdmin},
			{ID: 2, Name: "Bob", Role: RoleUser},
		},
	}
}

func (r *memoryRepo) GetAll() []User {
	return r.users
}
