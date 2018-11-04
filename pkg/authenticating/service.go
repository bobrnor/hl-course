package authenticating

import "fmt"

var ErrUserNotFound = fmt.Errorf("user not found")

type Service interface {
	Authenticate(token string) (id int, err error)
}

type Repository interface {
	FindUser(token string) (*User, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) Authenticate(token string) (id int, err error) {
	user, err := s.repo.FindUser(token)
	if err != nil {
		return 0, err
	}

	return user.ID, nil
}
