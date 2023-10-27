package value_objects

import (
	"testing"
)

func TestNewEmail_ValidEmail(t *testing.T) {
	_, err := NewEmail("test@example.com")
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestNewEmail_InvalidEmail_NoAtSymbol(t *testing.T) {
	_, err := NewEmail("testexample.com")
	if err == nil {
		t.Error("Expected an error for invalid email, but got none")
	}
}

func TestNewEmail_InvalidEmail_NoDomain(t *testing.T) {
	_, err := NewEmail("test@")
	if err == nil {
		t.Error("Expected an error for invalid email, but got none")
	}
}

func TestNewEmail_InvalidEmail_SpecialCharacters(t *testing.T) {
	_, err := NewEmail("test@exa$mple.com")
	if err == nil {
		t.Error("Expected an error for invalid email, but got none")
	}
}

// EmailのValueメソッドのテスト
func TestEmail_Value(t *testing.T) {
	email, err := NewEmail("test@example.com")
	if err != nil {
		t.Fatalf("Failed to create email: %v", err)
	}
	if email.Value() != "test@example.com" {
		t.Errorf("Expected email value to be 'test@example.com', but got: %v", email.Value())
	}
}
