package value_objects

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
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

// PasswordのValueメソッドのテスト
func TestPassword_Value(t *testing.T) {
	password, err := NewPassword("Aa1234567890!")
	if err != nil {
		t.Fatalf("Failed to create password: %v", err)
	}
	if len(password.Value()) == 0 {
		t.Error("Expected hashed password value, but got empty string")
	}
}

// ComparePassword関数のテスト
func TestComparePassword(t *testing.T) {
	password, err := NewPassword("Aa1234567890!")
	if err != nil {
		t.Fatalf("Failed to create password: %v", err)
	}
	isMatch, err := ComparePassword(password.Value(), "Aa1234567890!")
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if !isMatch {
		t.Error("Expected password to match, but it didn't")
	}
}

// 異常系: パスワードの長さが短い
func TestNewPassword_ShortLength(t *testing.T) {
	_, err := NewPassword("Aa1!")
	if err == nil {
		t.Error("Expected an error for short password length, but got none")
	}
}

// 異常系: パスワードのハッシュ化に失敗
func TestHashPassword_Error(t *testing.T) {
	// bcryptのcostを非常に高く設定してエラーを強制的に発生させる
	_, err := bcrypt.GenerateFromPassword([]byte("Aa1234567890!"), 100)
	if err == nil {
		t.Error("Expected an error during password hashing, but got none")
	}
}

// 正常系: 限定的な特殊文字を含む有効なパスワード
func TestIsValidPassword_LimitedSpecialCharacters(t *testing.T) {
    isValid := isValidPassword("Aa1234!@#$%^&*")
    if !isValid {
        t.Error("Expected password with limited special characters to be valid, but it wasn't")
    }
}

// 異常系: ( を含むパスワード
func TestIsValidPassword_WithOpenParenthesis(t *testing.T) {
    isValid := isValidPassword("Aa1234(")
    if isValid {
        t.Error("Expected password with '(' to be invalid, but it was valid")
    }
}

// 異常系: ) を含むパスワード
func TestIsValidPassword_WithCloseParenthesis(t *testing.T) {
    isValid := isValidPassword("Aa1234)")
    if isValid {
        t.Error("Expected password with ')' to be invalid, but it was valid")
    }
}

// 異常系: - を含むパスワード
func TestIsValidPassword_WithHyphen(t *testing.T) {
    isValid := isValidPassword("Aa1234-")
    if isValid {
        t.Error("Expected password with '-' to be invalid, but it was valid")
    }
}

// 異常系: _ を含むパスワード
func TestIsValidPassword_WithUnderscore(t *testing.T) {
    isValid := isValidPassword("Aa1234_")
    if isValid {
        t.Error("Expected password with '_' to be invalid, but it was valid")
    }
}


