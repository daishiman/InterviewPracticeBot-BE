package userusecase

import (
	"InterviewPracticeBot-BE/internal/domain/aggregates"
	"InterviewPracticeBot-BE/internal/domain/entities"
	"InterviewPracticeBot-BE/internal/domain/value_objects"
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

// UserRepositoryインタフェースのメソッドをモックとして実装
func (m *MockUserRepository) FindByID(id string) (*entities.UserPrivate, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.UserPrivate), args.Error(1)
}

func (m *MockUserRepository) FindByEmail(email string) (*entities.UserPrivate, error) {
	args := m.Called(email)
	return args.Get(0).(*entities.UserPrivate), args.Error(1)
}

func (m *MockUserRepository) Save(user *entities.UserPrivate) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestUserUsecase_Register(t *testing.T) {
	mockRepo := new(MockUserRepository)
	usecase := NewUserUsecase(mockRepo, aggregates.NewUserFactory())

	// 正常系
	t.Run("success", func(t *testing.T) {
		mockRepo.On("Save", mock.AnythingOfType("*entities.UserPrivate")).Return(nil)

		err := usecase.Register("test@example.com", "Password123!")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		mockRepo.AssertExpectations(t)
	})

	// 異常系t.Run("email already exists", func(t *testing.T) {
	t.Run("email already exists", func(t *testing.T) {
		// 既存のユーザーをシミュレートするためFindByEmailをモックします
		email, err := value_objects.NewEmail("test@example.com")
		if err != nil {
			t.Fatal(err)
		}
		password, err := value_objects.NewPassword("Password123!")
		if err != nil {
			t.Fatal(err)
		}
		mockUser := &entities.UserPrivate{Email: *email, Password: *password}
		mockRepo.On("FindByEmail", "test@example.com").Return(mockUser, nil)

		// このテストケースではSaveが呼び出されるべきではないため、Saveメソッドのモックは不要です

		err = usecase.Register("test@example.com", "Password123!")
		if err == nil || err.Error() != "user already exists" {
			t.Errorf("expected an error 'user already exists' but got: %v", err)
		}

		mockRepo.AssertExpectations(t)
	})

	t.Run("invalid password requirements", func(t *testing.T) {
		// このテストは、value objectのコンストラクタがエラーを返すことに依存しています
		err := usecase.Register("test@example.com", "password")
		if err == nil {
			t.Errorf("expected an error but got none")
		}
	})
}
