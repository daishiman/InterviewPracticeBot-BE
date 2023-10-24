package aggregates

import (
	"testing"
)

func TestUserFactory_CreateUser(t *testing.T) {
	factory := NewUserFactory()

	t.Run("正常系", func(t *testing.T) {
		user, err := factory.CreateUser("test@example.com", "Password123!")
		if err != nil {
			t.Fatal(err)
		}
		if user.Email.Value() != "test@example.com" {
			t.Error("Expected email to be test@example.com, got ", user.Email.Value())
		}
	})

	t.Run("異常系: 不正なメールアドレス", func(t *testing.T) {
		_, err := factory.CreateUser("invalid_email", "password")
		if err == nil {
			t.Error("Expected error for invalid email, got nil")
		}
	})

	t.Run("異常系: 不正なパスワード", func(t *testing.T) {
		_, err := factory.CreateUser("test@example.com", "1234")
		if err == nil {
			t.Error("Expected error for invalid password, got nil")
		}
	})

	t.Run("境界値: メールアドレスの長さが最大値", func(t *testing.T) {
		email := "a"
		for i := 0; i < 254; i++ {
			email += "a"
		}
		email += "@example.com"
		_, err := factory.CreateUser(email, "Password123!")
		if err != nil {
			t.Error("Expected no error for max length email, got ", err)
		}
	})

	t.Run("境界値: メールアドレスの長さが最大値を超える", func(t *testing.T) {
		email := "a"
		for i := 0; i < 255; i++ {
			email += "a"
		}
		email += "@example.com"
		_, err := factory.CreateUser(email, "password")
		if err == nil {
			t.Error("Expected error for email length exceeding max, got nil")
		}
	})

	t.Run("値がnil", func(t *testing.T) {
		_, err := factory.CreateUser("", "")
		if err == nil {
			t.Error("Expected error for empty email and password, got nil")
		}
	})
}
