package registration

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/docker/distribution/uuid"
)

var ErrUserAlreadyExists = fmt.Errorf("user already exists")
var ErrProfileAlreadyExists = fmt.Errorf("profile already exists")

type Service interface {
	RegisterUser() (User, error)
}

type Repository interface {
	CreateUser(u User) (id int, err error)
	CreateProfile(userID int) error
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

	id, err := s.repo.CreateUser(user)
	if err != nil {
		return User{}, err
	}

	if err := s.repo.CreateProfile(id); err != nil {
		if errors.Cause(err) != ErrProfileAlreadyExists {
			// TODO: delete previously created user
			return User{}, err
		}
	}

	return user, nil
}
