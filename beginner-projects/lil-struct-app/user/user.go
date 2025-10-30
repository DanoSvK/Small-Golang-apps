package user

import "time"

type User struct {
	FirstName string
	LastName  string
	BirthDate string
	UpdatedAt time.Time
}
