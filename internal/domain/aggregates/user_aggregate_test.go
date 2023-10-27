package aggregates

import (
	"InterviewPracticeBot-BE/internal/domain/entities"
	"InterviewPracticeBot-BE/internal/domain/repositories"
	"testing"
)

type mockUserRepository struct {
	repositories.IUserPrivateRepository
}

func (m *mockUserRepository) Save(user *entities.UserPrivate) error {
	return nil
}

func TestUserAggregate_NewUserAggregate(t *testing.T) {
	factory := NewUserFactory()

	t.Run("正常系", func(t *testing.T) {
		_, err := NewUserAggregate(factory, "test@example.com", "Password123!")
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("異常系: 不正なメールアドレス", func(t *testing.T) {
		_, err := NewUserAggregate(factory, "invalid_email", "Password123!")
		if err == nil {
			t.Error("Expected error for invalid email, got nil")
		}
	})

	t.Run("異常系: 不正なパスワード", func(t *testing.T) {
		_, err := NewUserAggregate(factory, "test@example.com", "pass")
		if err == nil {
			t.Error("Expected error for invalid password, got nil")
		}
	})
}

func TestUserAggregate_UpdatePassword(t *testing.T) {
	factory := NewUserFactory()
	userRepo := &mockUserRepository{}

	t.Run("正常系", func(t *testing.T) {
		aggregate, _ := NewUserAggregate(factory, "test@example.com", "Password123!")
		err := aggregate.UpdatePassword("Password123!", "newPassword123!", userRepo)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("異常系: 不正な旧パスワード", func(t *testing.T) {
		aggregate, _ := NewUserAggregate(factory, "test@example.com", "wrongPassword123!")
		err := aggregate.UpdatePassword("wrong", "newPassword123!", userRepo)
		if err == nil {
			t.Error("Expected error for incorrect old password, got nil")
		}
	})

	t.Run("異常系: 不正な新パスワード", func(t *testing.T) {
		aggregate, _ := NewUserAggregate(factory, "test@example.com", "Password123!")
		err := aggregate.UpdatePassword("Password123!", "", userRepo)
		if err == nil {
			t.Error("Expected error for invalid new password, got nil")
		}
	})
}
