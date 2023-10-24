package entities

import (
	"InterviewPracticeBot-BE/internal/domain/value_objects"
	"time"
)

type UserPrivate struct {
	ID           string
	Email        value_objects.Email
	Password     value_objects.Password
	Verification value_objects.Verification
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}

var _ IUserPrivate = (*UserPrivate)(nil)

func (u *UserPrivate) GetID() string {
	return u.ID
}

func (u *UserPrivate) GetEmail() value_objects.Email {
	return u.Email
}

func (u *UserPrivate) GetPassword() value_objects.Password {
	return u.Password
}

func (u *UserPrivate) GetVerification() value_objects.Verification {
	return u.Verification
}

func (u *UserPrivate) GetCreatedAt() time.Time {
	return u.CreatedAt
}

func (u *UserPrivate) GetUpdatedAt() time.Time {
	return u.UpdatedAt
}

func (u *UserPrivate) GetDeletedAt() *time.Time {
	return u.DeletedAt
}
