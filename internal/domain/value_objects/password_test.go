package value_objects

import (
	"testing"
)

// 正常系: パスワードのバリデーションが正しく動作するか
func TestNewPassword_ValidPassword(t *testing.T) {
	_, err := NewPassword("Aa1234567890!")
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

// 異常系: 小文字が含まれていない
func TestNewPassword_NoLowercase(t *testing.T) {
	_, err := NewPassword("A1234567890!")
	if err == nil {
		t.Error("Expected an error for missing lowercase, but got none")
	}
}

// 異常系: 大文字が含まれていない
func TestNewPassword_NoUppercase(t *testing.T) {
	_, err := NewPassword("a1234567890!")
	if err == nil {
		t.Error("Expected an error for missing uppercase, but got none")
	}
}

// 異常系: 数字が含まれていない
func TestNewPassword_NoDigit(t *testing.T) {
	_, err := NewPassword("Aa!!!!!!!!!!")
	if err == nil {
		t.Error("Expected an error for missing digit, but got none")
	}
}

// 異常系: 特殊文字が含まれていない
func TestNewPassword_NoSpecialCharacter(t *testing.T) {
	_, err := NewPassword("Aa1234567890")
	if err == nil {
		t.Error("Expected an error for missing special character, but got none")
	}
}

