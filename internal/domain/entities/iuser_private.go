package entities

import (
	"InterviewPracticeBot-BE/internal/domain/value_objects"
	"time"
)

type IUserPrivate interface {
	GetID() *value_objects.UUID
	GetEmail() value_objects.Email
	GetPassword() value_objects.Password
	GetVerification() value_objects.Verification
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	GetDeletedAt() *time.Time
}
