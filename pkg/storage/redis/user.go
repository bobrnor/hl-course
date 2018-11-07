package redis

import "encoding/json"

type User struct {
	ID    int
	Token string
}

func (u User) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}

func (u *User) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &u); err != nil {
		return err
	}
	return nil
}
