package aggregates

import (
	"InterviewPracticeBot-BE/internal/domain/entities"
	"InterviewPracticeBot-BE/internal/domain/repositories"
	"InterviewPracticeBot-BE/internal/domain/value_objects"
	"errors"
)

type UserAggregate struct {
	User *entities.UserPrivate
}

var _ IUserAggregate = (*UserAggregate)(nil)

func NewUserAggregate(factory *UserFactory, email, rawPassword string) (*UserAggregate, error) {
	user, err := factory.CreateUser(email, rawPassword)
	if err != nil {
		return nil, err
	}
	return &UserAggregate{
		User: user,
	}, nil
}

func (ua *UserAggregate) Register() (*entities.UserPrivate, error) {
	return ua.User, nil
}

func (ua *UserAggregate) UpdatePassword(oldPassword, NewPassword string, userRepo repositories.IUserPrivateRepository) error {
	isCorrect, err := value_objects.ComparePassword(ua.User.Password.Value(), oldPassword)
	if err != nil || !isCorrect {
		return errors.New("incorrect old password")
	}

	newPasswordObj, err := value_objects.NewPassword(NewPassword)
	if err != nil {
		return err
	}
	ua.User.Password = *newPasswordObj
	return userRepo.Save(ua.User)
}
