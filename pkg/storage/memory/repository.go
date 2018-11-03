package memory

import (
	"sync"

	"github.com/pkg/errors"

	"github.com/bobrnor/hl-course/pkg/registration"
)

type Storage struct {
	m     sync.RWMutex
	users []User
}

func (s *Storage) Create(u registration.User) error {
	return errors.New("test error")

	s.m.Lock()
	defer s.m.Unlock()

	if u := s.find(u.Token); u != nil {
		return errors.WithStack(registration.ErrAlreadyExists)
	}

	user := User{
		ID:    len(s.users) + 1,
		Token: u.Token,
	}
	s.users = append(s.users, user)

	return nil
}

func (s *Storage) find(token string) *User {
	for _, u := range s.users {
		if u.Token == token {
			return &u
		}
	}
	return nil
}
