package registration

import (
	"fmt"

	"github.com/docker/distribution/uuid"
)

var ErrAlreadyExists = fmt.Errorf("user already exists")

type Service interface {
	RegisterUser() (User, error)
}

type Repository interface {
	Create(u User) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) RegisterUser() (User, error) {
	token := uuid.Generate().String()
	user := User{
		Token: token,
	}

	if err := s.repo.Create(user); err != nil {
		return User{}, err
	}

	return user, nil
}
