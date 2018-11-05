package editing

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Profile struct {
	FirstName *string    `json:"first_name"`
	LastName  *string    `json:"last_name"`
	BirthDate *time.Time `json:"birth_date"`
	Status    *string    `json:"status"`
}

func (p Profile) Validate() error {
	return validation.ValidateStruct(
		&p,
		validation.Field(&p.FirstName, validation.NilOrNotEmpty, validation.Length(0, 80)),
		validation.Field(&p.LastName, validation.NilOrNotEmpty, validation.Length(0, 80)),
		validation.Field(&p.BirthDate, validation.Max(time.Now())),
		validation.Field(&p.Status, validation.Length(0, 255)),
	)
}
