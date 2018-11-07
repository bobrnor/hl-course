package redis

import (
	"encoding/json"
	"time"
)

type Profile struct {
	ID        int
	UserID    int
	FirstName string
	LastName  string
	BirthDate time.Time
	Status    string
}

func (p Profile) MarshalBinary() ([]byte, error) {
	return json.Marshal(p)
}

func (p *Profile) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &p); err != nil {
		return err
	}
	return nil
}
