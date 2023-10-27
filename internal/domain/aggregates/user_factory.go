package aggregates

import (
	"InterviewPracticeBot-BE/internal/domain/entities"
	"InterviewPracticeBot-BE/internal/domain/value_objects"
	"time"
)

type UserFactory struct{}

func NewUserFactory() *UserFactory {
	return &UserFactory{}
}

func (uf *UserFactory) CreateUser(emailVO value_objects.Email, passwordVO value_objects.Password) (*entities.UserPrivate, error) {
	uuidVO, err := uf.createUUID()
	if err != nil {
		return nil, err
	}

	verificationVO, err := uf.createVerification()
	if err != nil {
		return nil, err
	}

	user := &entities.UserPrivate{
		ID:           uuidVO,
		Email:        emailVO,
		Password:     passwordVO,
		Verification: *verificationVO,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	return user, nil
}

func (uf *UserFactory) createUUID() (*value_objects.UUID, error) {
	uuid, err := value_objects.NewUUID()
	if err != nil {
		return nil, err
	}
	return uuid, nil
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
