package user

import "time"

type User struct {
	ID             uint
	Name           string
	Email          string
	Occupation     string
	PasswordHash   string
	AvatarFileName string
	Role           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
