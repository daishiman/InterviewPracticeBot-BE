package userusecase

import (
	"InterviewPracticeBot-BE/internal/domain/aggregates"
	"InterviewPracticeBot-BE/internal/domain/repositories"
)

type UserUsecase struct {
	repo    repositories.UserRepository
	factory *aggregates.UserFactory
}

// ユースケースの新しいインスタンスを作成するコンストラクタ
func NewUserUsecase(repo repositories.UserRepository, factory *aggregates.UserFactory) *UserUsecase {
	return &UserUsecase{
		repo:    repo,
		factory: factory,
	}
}

// ユーザー登録のロジック
func (uu *UserUsecase) Register(email, password string) error {
	user, err := uu.factory.CreateUser(email, password)
	if err != nil {
		return err
	}
	return uu.repo.Save(user)
}

// パスワード更新のロジック
func (uu *UserUsecase) UpdatePassword(email, oldPassword, newPassword string) error {
	user, err := uu.repo.FindByEmail(email)
	if err != nil {
		return err
	}
	aggregatedUser := &aggregates.UserAggregate{User: user}
	return aggregatedUser.UpdatePassword(oldPassword, newPassword, uu.repo)
}
