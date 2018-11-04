package editing

import "time"

type Profile struct {
	FirstName *string    `json:"first_name,omitempty",validate:"min=1,max=80"`
	LastName  *string    `json:"last_name,omitempty",validate:"min=1,max=80"`
	BirthDate *time.Time `json:"birth_date,omitempty",validate:"lt"`
	Status    *string    `json:"status,omitempty",validate:"max=255"`
}
