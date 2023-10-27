package repositories

import (
	"InterviewPracticeBot-BE/internal/domain/entities"
	"InterviewPracticeBot-BE/internal/domain/value_objects"
	"testing"
	"time"
)

func createTestUser(emailValue string, passwordValue string) (*entities.UserPrivate, error) {
	uuidObj, err := value_objects.NewUUID()
	if err != nil {
		return nil, err
	}

	emailObj, err := value_objects.NewEmail(emailValue)
	if err != nil {
		return nil, err
	}

	passwordObj, err := value_objects.NewPassword(passwordValue)
	if err != nil {
		return nil, err
	}

	return &entities.UserPrivate{
		ID:        uuidObj,
		Email:     *emailObj,
		Password:  *passwordObj,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

// Saveメソッドの正常系テスト
func TestInMemoryUserRepository_Save_NewUser(t *testing.T) {
	repo := NewInMemoryUserRepository()

	user, err := createTestUser("test@example.com", "Password1234!")
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	err = repo.Save(user)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

// Saveメソッドの異条系テスト: 重複ユーザー
func TestInMemoryUserRepository_Save_DuplicateUser(t *testing.T) {
	repo := NewInMemoryUserRepository()

	user, err := createTestUser("test@example.com", "Password1234!")
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	repo.Save(user)
	err = repo.Save(user)
	if err == nil {
		t.Error("Expected an error for dupicate user, but got none")
	}
}

// Saveメソッドの異常系テスト: 無効なデータ
func TestInMemoryUserRepository_Save_InvalidData(t *testing.T) {
	user, err := createTestUser("", "")
	if user != nil || err == nil {
		t.Error("Expected an error for invalid data, but got none")
	}
}

// FindByIDメソッドの正常系テスト
func TestInMemoryUserRepository_FindByID_ExisingUser(t *testing.T) {
	repo := NewInMemoryUserRepository()
	user, err := createTestUser("test@example.com", "Password1234!")
	if err != nil {
		t.Fatalf("faild to create user: %v", err)
	}
	repo.Save(user)
	foundUser, err := repo.FindByID(user.ID.String())
	if err != nil || foundUser.ID.String() != "1" {
		t.Errorf("Expeted to find user with ID 1, but got: %v, err: %v", foundUser, err)
	}
}

// FindByIDメソッドの異常系のテスト
func TestInMemoryUserRepository_FindByID_NonExistentUser(t *testing.T) {
	repo := NewInMemoryUserRepository()
	_, err := repo.FindByID("2")
	if err == nil {
		t.Errorf("Expeted an error for non-existent user, but got none")
	}
}

// FindByIDメソッドの異常系テスト: 無効なID
func TestInMemoryUserRepository_FindByID_InvalidID(t *testing.T) {
	repo := NewInMemoryUserRepository()
	_, err := repo.FindByID("")
	if err == nil {
		t.Errorf("Expetced an error fo invalid ID, but got none")
	}
}

// FindByEmailメソッドの正常系テスト
func TestInMemoryUserRepository_FindByEmail_ExistingEmail(t *testing.T) {
	repo := NewInMemoryUserRepository()
	user, err := createTestUser("test@example.com", "Password1234!")
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	repo.Save(user)
	foundUser, err := repo.FindByEmail("test@example.com")
	if err != nil || foundUser.Email.Value() != "test@example.com" {
		t.Errorf("Expetced to find user with email test@example.com, but got: %v, error: %v", foundUser, err)
	}
}

// FindByEmailメソッドの異常系テスト
func TestInMemoryUserRepository_FindByEmail_NonExistentEmail(t *testing.T) {
	repo := NewInMemoryUserRepository()
	_, err := repo.FindByEmail("notfound@example.com")
	if err == nil {
		t.Errorf("Expetced an error fo non-existent email, but got none")
	}
}

// FindByEmailメソッドの異常系テスト: 無効なメールアドレス
func TestInMemoryUserRepository_FindByEmail_InvalidEmail(t *testing.T) {
	repo := NewInMemoryUserRepository()
	_, err := repo.FindByEmail("")
	if err == nil {
		t.Error("Expected an error for invalid email, but got none")
	}
}

// Deleteメソッドの正常系テスト
func TestInMemoryUserRepository_Delete_ExistingUser(t *testing.T) {
	repo := NewInMemoryUserRepository()
	user, err := createTestUser("test@example.com", "Password1234!")
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}
	repo.Save(user)
	err = repo.Delete("1")
	if err != nil {
		t.Errorf("Expected no error when deleting, but got: %v", err)
	}

	deletedUser, err := repo.FindByID("1")
	if err == nil {
		t.Errorf("Expected an error when finding deleted user, but got none")
	}
	if deletedUser != nil {
		t.Errorf("User should have been deleted, but was found in the repository")
	}
}

// Deleteメソッドの異常系テスト
func TestInMemoryUserRepository_Delete_NonExistentUser(t *testing.T) {
	repo := NewInMemoryUserRepository()
	err := repo.Delete("1")
	if err == nil {
		t.Error("Expected an error for non-existent user, but got none")
	}

	nonExistentUser, err := repo.FindByID("1") // 仮にFindというメソッドが存在するとして
	if err == nil {
		t.Errorf("Expected an error when finding non-existent user, but got none")
	}
	if nonExistentUser != nil {
		t.Errorf("User should not exist, but was found in the repository")
	}
}

// Deleteメソッドの異常系テスト: 無効なID
func TestInMemoryUserRepository_Delete_InvalidID(t *testing.T) {
	repo := NewInMemoryUserRepository()
	err := repo.Delete("")
	if err == nil {
		t.Error("Expected an error for invalid ID, but got none")
	}
}
