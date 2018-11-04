package memory

import (
	"sync"
	"time"

	"github.com/pkg/errors"

	"github.com/bobrnor/hl-course/pkg/registration"
)

type Storage struct {
	m        sync.RWMutex
	users    []User
	profiles []Profile
}

func (s *Storage) Create(u registration.User) (id int, err error) {
	s.m.Lock()
	defer s.m.Unlock()

	if u := s.findUser(u.Token); u != nil {
		return 0, errors.WithStack(registration.ErrUserAlreadyExists)
	}

	user := User{
		ID:    len(s.users) + 1,
		Token: u.Token,
	}
	s.users = append(s.users, user)

	return user.ID, nil
}

func (s *Storage) CreateProfile(userID int) error {
	s.m.Lock()
	defer s.m.Unlock()

	if p := s.findProfile(userID); p != nil {
		return errors.WithStack(registration.ErrProfileAlreadyExists)
	}

	profile := Profile{
		UserID:    userID,
		FirstName: "",
		LastName:  "",
		BirthDate: time.Time{},
		Status:    "",
	}
	s.profiles = append(s.profiles, profile)

	return nil
}

func (s *Storage) findUser(token string) *User {
	for _, u := range s.users {
		if u.Token == token {
			return &u
		}
	}
	return nil
}

func (s *Storage) findProfile(userID int) *Profile {
	for _, p := range s.profiles {
		if p.UserID == userID {
			return &p
		}
	}
	return nil
}
