package redis

import (
	"strconv"
	"time"

	"github.com/bobrnor/hl-course/pkg/authenticating"
	"github.com/bobrnor/hl-course/pkg/editing"
	"github.com/bobrnor/hl-course/pkg/registration"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

const (
	LastUserIDKey    = "last-id"
	LastProfileIDKey = "last-id"

	NoExpiration = 0
)

type Storage struct {
	client *redis.Client
}

func NewStorage(c *redis.Client) *Storage {
	return &Storage{
		client: c,
	}
}

func (s *Storage) CreateUser(u registration.User) (id int, err error) {
	nextID, err := s.client.Incr(LastUserIDKey).Result()
	if err != nil {
		return 0, errors.WithStack(err)
	}

	user := User{
		ID:    int(nextID),
		Token: u.Token,
	}

	ok, err := s.client.SetNX(user.Token, user, NoExpiration).Result()
	if err != nil {
		return 0, errors.WithStack(err)
	}

	if !ok {
		return 0, errors.WithStack(registration.ErrUserAlreadyExists)
	}

	return user.ID, nil
}

func (s *Storage) CreateProfile(userID int) error {
	nextID, err := s.client.Incr(LastProfileIDKey).Result()
	if err != nil {
		return errors.WithStack(err)
	}

	profile := Profile{
		ID:        int(nextID),
		UserID:    userID,
		FirstName: "",
		LastName:  "",
		BirthDate: time.Time{},
		Status:    "",
	}

	userIDKey := strconv.Itoa(profile.UserID)

	ok, err := s.client.SetNX(userIDKey, profile, NoExpiration).Result()
	if err != nil {
		return errors.WithStack(err)
	}

	if !ok {
		return errors.WithStack(registration.ErrProfileAlreadyExists)
	}

	return nil
}

func (s *Storage) FindUser(token string) (*authenticating.User, error) {
	var u User
	if err := s.client.Get(token).Scan(&u); err != nil {
		return nil, errors.WithStack(err)
	}

	authUser := authenticating.User{
		ID:    u.ID,
		Token: u.Token,
	}

	return &authUser, nil
}

func (s *Storage) UpdateProfile(userID int, p editing.Profile) error {
	var profile Profile

	userIDKey := strconv.Itoa(userID)
	if err := s.client.Get(userIDKey).Scan(&profile); err != nil {
		return errors.WithStack(err)
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

	if _, err := s.client.Set(userIDKey, profile, NoExpiration).Result(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
