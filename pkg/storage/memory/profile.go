package memory

import "time"

type Profile struct {
	ID        int
	UserID    int
	FirstName string
	LastName  string
	BirthDate time.Time
	Status    string
}
