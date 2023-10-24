package aggregates

import (
	"InterviewPracticeBot-BE/internal/domain/entities"
	"InterviewPracticeBot-BE/internal/domain/utilities"
	"InterviewPracticeBot-BE/internal/domain/value_objects"
	"time"
)

type UserFactory struct{}

func NewUserFactory() *UserFactory {
	return &UserFactory{}
}

func (uf *UserFactory) CreateUser(email, rawPassword string) (*entities.UserPrivate, error) {
	emailVO, err := uf.createEmail(email)
	if err != nil {
		return nil, err
	}

	passwordVO, err := uf.createPassword(rawPassword)
	if err != nil {
		return nil, err
	}

	verificationVO, err := uf.createVerification()
	if err != nil {
		return nil, err
	}

	user := &entities.UserPrivate{
		ID:           utilities.GenerateUUID(),
		Email:        *emailVO,
		Password:     *passwordVO,
		Verification: *verificationVO,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	return user, nil
}

func (uf *UserFactory) createEmail(email string) (*value_objects.Email, error) {
	return value_objects.NewEmail(email)
}

func (uf *UserFactory) createPassword(rawPassword string) (*value_objects.Password, error) {
	return value_objects.NewPassword(rawPassword)
}

func (uf *UserFactory) createVerification() (*value_objects.Verification, error) {
	return value_objects.NewVerification()
}
