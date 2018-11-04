package memory

import "time"

type Profile struct {
	UserID    int
	FirstName string
	LastName  string
	BirthDate time.Time
	Status    string
}
