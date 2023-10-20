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

func NewUserAggregate(factory *UserFactory, email, rawPassword string) (*UserAggregate, error) {
  user, err := factory.CreateUser(email, rawPassword)
  if err != nil {
    return nil, err
  }
  return &UserAggregate{
    User: user,
  }, nil
}

func (ua *UserAggregate) Register(userRepo repositories.UserRepository) error {
	// todo: 1. ユーザーがすでに存在するか確認
	existingUser, err := userRepo.FindByEmail(ua.User.Email.Value())
	if err != nil && err != repositories.ErrUserNotFound {
		return err
	}
	if existingUser != nil {
		return errors.New("user already exists")
	}

	// todo: 2. 新しいユーザーをデータベースに保存
	err = userRepo.Save(ua.User)
	if err != nil {
		return err
	}

	return nil
}

func (ua *UserAggregate) UpdatePassword(oldPassword, NewPassword string, userRepo repositories.UserRepository) error {
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
