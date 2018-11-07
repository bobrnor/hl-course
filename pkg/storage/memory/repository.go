package memory

import (
	"sync"
	"time"

	"github.com/bobrnor/hl-course/pkg/authenticating"

	"github.com/bobrnor/hl-course/pkg/editing"

	"github.com/pkg/errors"

	"github.com/bobrnor/hl-course/pkg/registration"
)

type Storage struct {
	m        sync.RWMutex
	users    []User
	profiles []Profile
}

func (s *Storage) CreateUser(u registration.User) (id int, err error) {
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

func (s *Storage) FindUser(token string) (*authenticating.User, error) {
	s.m.RLock()
	defer s.m.RUnlock()

	user := s.findUser(token)
	if user == nil {
		return nil, errors.WithStack(authenticating.ErrUserNotFound)
	}

	authUser := authenticating.User{
		ID:    user.ID,
		Token: user.Token,
	}

	return &authUser, nil
}

func (s *Storage) CreateProfile(userID int) error {
	s.m.Lock()
	defer s.m.Unlock()

	if p := s.findProfile(userID); p != nil {
		return errors.WithStack(registration.ErrProfileAlreadyExists)
	}

	profile := Profile{
		ID:        len(s.profiles) + 1,
		UserID:    userID,
		FirstName: "",
		LastName:  "",
		BirthDate: time.Time{},
		Status:    "",
	}
	s.profiles = append(s.profiles, profile)

	return nil
}

func (s *Storage) UpdateProfile(userID int, p editing.Profile) error {
	s.m.Lock()
	defer s.m.Unlock()

	profile := s.findProfile(userID)
	if profile == nil {
		return errors.WithStack(editing.ErrProfileNotFound)
	}

	if p.FirstName != nil {
		profile.FirstName = *p.FirstName
	}

	if p.LastName != nil {
		profile.LastName = *p.LastName
	}

	if p.BirthDate != nil {
		profile.BirthDate = time.Unix(*p.BirthDate, 0)
	}

	if p.Status != nil {
		profile.Status = *p.Status
	}

	return s.replaceProfile(*profile)
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

func (s *Storage) replaceProfile(p Profile) error {
	if p.ID <= 0 || p.ID > len(s.profiles) {
		return errors.Errorf("profile id (%d) is out of range [0..%d]", p.ID, len(s.profiles))
	}

	s.profiles[p.ID-1] = p

	return nil
}
