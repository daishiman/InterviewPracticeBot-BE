package aggregates

import (
	"InterviewPracticeBot-BE/internal/domain/value_objects"
	"testing"
	"time"
)

func TestUserFactory_CreateUser(t *testing.T) {
	factory := NewUserFactory()

	t.Run("正常系", func(t *testing.T) {
		emailVO, _ := value_objects.NewEmail("test@example.com")
		passwordVO, _ := value_objects.NewPassword("Password123!")
		user, err := factory.CreateUser(*emailVO, *passwordVO)

		if err != nil {
			t.Fatal(err)
		}
		if user.Email.Value() != "test@example.com" {
			t.Error("Expected email to be test@example.com, got ", user.Email.Value())
		}
	})

	t.Run("異常系: 不正なメールアドレス", func(t *testing.T) {
		emailVO, _ := value_objects.NewEmail("invalid_email")
		passwordVO, _ := value_objects.NewPassword("password")
		_, err := factory.CreateUser(*emailVO, *passwordVO)
		if err == nil {
			t.Error("Expected error for invalid email, got nil")
		}
	})

	t.Run("異常系: 不正なパスワード", func(t *testing.T) {
		emailVO, _ := value_objects.NewEmail("test@example.com")
		passwordVO, _ := value_objects.NewPassword("1234")
		_, err := factory.CreateUser(*emailVO, *passwordVO)
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
		emailVO, _ := value_objects.NewEmail(email)
		passwordVO, _ := value_objects.NewPassword("Password123!")
		_, err := factory.CreateUser(*emailVO, *passwordVO)
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
		emailVO, _ := value_objects.NewEmail(email)
		passwordVO, _ := value_objects.NewPassword("password")
		_, err := factory.CreateUser(*emailVO, *passwordVO)
		if err == nil {
			t.Error("Expected error for email length exceeding max, got nil")
		}
	})

	t.Run("値がnil", func(t *testing.T) {
		emailVO, _ := value_objects.NewEmail("")
		passwordVO, _ := value_objects.NewPassword("")
		_, err := factory.CreateUser(*emailVO, *passwordVO)
		if err == nil {
			t.Error("Expected error for empty email and password, got nil")
		}
	})

	// 正常系: createUUID
	t.Run("createUUID generates valid UUID", func(t *testing.T) {
		uuid, err := factory.createUUID()
		if err != nil {
			t.Fatal(err)
		}
		if _, err := value_objects.UUIDFromString(uuid.String()); err != nil {
			t.Errorf("Expected valid UUID format, but got: %v", err)
		}
	})

	// 正常系: createEmail
	t.Run("createEmail creates valid email", func(t *testing.T) {
		emailStr := "test@example.com"
		email, err := factory.createEmail(emailStr)
		if err != nil {
			t.Fatal(err)
		}
		if email.Value() != emailStr {
			t.Errorf("Expected email to be %s, but got %s", emailStr, email.Value())
		}
	})

	// 異常系: createEmail with invalid email
	t.Run("createEmail with invalid email", func(t *testing.T) {
		_, err := factory.createEmail("invalid_email")
		if err == nil {
			t.Error("Expected error for invalid email format, but got none")
		}
	})

	// 正常系: createPassword
	t.Run("createPassword creates valid password", func(t *testing.T) {
		passwordStr := "Password123!"
		password, err := factory.createPassword(passwordStr)
		if err != nil {
			t.Fatal(err)
		}
		if password.Value() != passwordStr {
			t.Errorf("Expected password to be %s, but got %s", passwordStr, password.Value())
		}
	})

	// 異常系: createPassword with short password
	t.Run("createPassword with short password", func(t *testing.T) {
		_, err := factory.createPassword("1234")
		if err == nil {
			t.Error("Expected error for short password, but got none")
		}
	})

	// 正常系: createVerification
	t.Run("createVerification", func(t *testing.T) {
		verification, err := factory.createVerification()
		if err != nil {
			t.Fatal(err)
		}

		// トークンが空でないことを確認
		if verification.Token == "" {
			t.Error("Expected non-empty token, but got empty string")
		}

		// トークンの有効期限が正確に1時間後であることを確認
		expectedExpiry := time.Now().Add(1 * time.Hour)
		if verification.Expiry.Sub(expectedExpiry) > time.Minute {
			t.Errorf("Expected token expiry to be roughly 1 hour from now, but got: %v", verification.Expiry)
		}

		// 生成直後のトークンが有効であることを確認
		if !verification.IsValid() {
			t.Error("Expected the token to be valid right after creation, but it was not")
		}
	})

}
