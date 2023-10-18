package entities

import (
	"InterviewPracticeBot-BE/internal/domain/value_objects"
	"time"
)

type UserPrivate struct {
	ID        string
	Email     value_objects.Email
	Password  value_objects.Password
	CreatedAt time.Time
	UpdatedAt time.Time
}
