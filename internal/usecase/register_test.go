package usecase

import "testing"

func TestRegisterUser_Success(t *testing.T) {
	mockRepo := NewMockUserRepository()
	uc := NewRegisterUsecase(mockRepo)

	err := uc.Register("test@example.com", "ValidPassword123!")
	if err != nil {
		t.Fatalf("expected no error, but got %v", err)
	}
}

func TestRegisterUser_EmailAlreadyExists(t *testing.T) {
	mockRepo := NewMockUserRepository()
	mockRepo.ExistsEmailFunc = func(email string) bool {
		return email == "existing@example.com"
	}

	uc := NewRegisterUsecase(mockRepo)

	err := uc.Register("existing@example.com", "ValidPassword123!")

	if err == nil {
		t.Fatal("expected an error, but got none")
	}

	if err != ErrEmailAlreadyExists {
		t.Fatalf("expected %v, but got %v", ErrEmailAlreadyExists, err)
	}
}

func TestRegisterUser_InvalidEmail(t *testing.T) {
	mockRepo := NewMockUserRepository()
	uc := NewRegisterUsecase(mockRepo)

	err := uc.Register("invalid-email", "ValidPassword123!")

	if err == nil {
		t.Fatal("expected an error, but got none")
	}

	if err != ErrInvalidEmail {
		t.Fatalf("expected %v, but got %v", ErrInvalidEmail, err)
	}
}
