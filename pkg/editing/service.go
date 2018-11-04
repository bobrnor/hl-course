package editing

import (
	"fmt"

	validator "gopkg.in/go-playground/validator.v9"

	"github.com/pkg/errors"
)

var ErrProfileNotFound = fmt.Errorf("profile not found")

type Service interface {
	EditProfile(userID int, p Profile) error
}

type Repository interface {
	UpdateProfile(userID int, p Profile) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) EditProfile(userID int, p Profile) error {
	validate := validator.New()

	if err := validate.Struct(p); err != nil {
		return errors.WithStack(err)
	}

	if err := s.repo.UpdateProfile(userID, p); err != nil {
		return err
	}

	return nil
}
