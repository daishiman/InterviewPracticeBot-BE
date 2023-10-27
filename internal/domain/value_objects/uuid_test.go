package value_objects

import (
	"strings"
	"testing"
)

// 正常系: UUIDの生成が正しく動作するか
func TestNewUUID(t *testing.T) {
	uuid, err := NewUUID()
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if uuid.String() == "" {
		t.Error("Expected non-empty UUID, but got empty string")
	}
}

// 正常系: UUIDの文字列からの変換が正しく動作するか
func TestUUIDFromString_ValidUUID(t *testing.T) {
	originalUUID, err := NewUUID()
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	convertedUUID, err := UUIDFromString(originalUUID.String())
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if !convertedUUID.Equals(originalUUID) {
		t.Error("Expected UUIDs to be equal, but they were not")
	}
}

// 異常系: 不正なUUID形式の文字列
func TestUUIDFromString_InvalidUUID(t *testing.T) {
	_, err := UUIDFromString("invalid-uuid-string")
	if err == nil {
		t.Error("Expected an error for invalid UUID string, but got none")
	}
}

// 異常系: 空の文字列
func TestUUIDFromString_EmptyString(t *testing.T) {
	_, err := UUIDFromString("")
	if err == nil {
		t.Error("Expected an error for empty string, but got none")
	}
}

// 異常系: 長すぎるUUID形式の文字列
func TestUUIDFromString_TooLongUUID(t *testing.T) {
	_, err := UUIDFromString(strings.Repeat("a", 200))
	if err == nil {
		t.Error("Expected an error for too long UUID string, but got none")
	}
}

// 異常系: 短すぎるUUID形式の文字列
func TestUUIDFromString_TooShortUUID(t *testing.T) {
	_, err := UUIDFromString("a")
	if err == nil {
		t.Error("Expected an error for too short UUID string, but got none")
	}
}

// UUIDのStringメソッドのテスト
func TestUUID_String(t *testing.T) {
	uuid, err := NewUUID()
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if len(uuid.String()) != 36 { // UUIDの形式は常に36文字
		t.Errorf("Expected UUID string length to be 36, but got: %d", len(uuid.String()))
	}
}

// UUIDのEqualsメソッドのテスト
func TestUUID_Equals(t *testing.T) {
	uuid1, err := NewUUID()
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	uuid2, err := NewUUID()
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if uuid1.Equals(uuid2) {
		t.Error("Expected different UUIDs to be not equal, but they were")
	}

	if !uuid1.Equals(uuid1) {
		t.Error("Expected the same UUID to be equal, but it was not")
	}
}
