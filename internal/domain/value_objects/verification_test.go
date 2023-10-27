package value_objects_test

import (
	"encoding/base64"
	"testing"
	"time"

	"InterviewPracticeBot-BE/internal/domain/value_objects"

	"github.com/stretchr/testify/assert"
)

// TestNewVerification_NoError は NewVerification 関数がエラーを返さない。
func TestNewVerification_NoError(t *testing.T) {
	_, err := value_objects.NewVerification()
	assert.NoError(t, err)
}

// TestNewVerification_TokenLength は生成されたトークンが32バイトの長さである。
func TestNewVerification_TokenLength(t *testing.T) {
	verification, _ := value_objects.NewVerification()
	decodedToken, _ := base64.URLEncoding.DecodeString(verification.Token)
	assert.Equal(t, 32, len(decodedToken))
}

// TestNewVerification_TokenEncoding は生成されたトークンがbase64エンコードされている。
func TestNewVerification_TokenEncoding(t *testing.T) {
	verification, _ := value_objects.NewVerification()
	_, err := base64.URLEncoding.DecodeString(verification.Token)
	assert.NoError(t, err)
}

// TestNewVerification_ExpireTime は有効期限が現在時刻から1時間後である。
func TestNewVerification_ExpireTime(t *testing.T) {
	verification, _ := value_objects.NewVerification()
	assert.WithinDuration(t, time.Now().Add(1*time.Hour), verification.Expiry, 1*time.Second)
}

// TestIsValid_True は有効期限が現在時刻よりも後の場合、IsValid メソッドが true を返す。
func TestIsValid_True(t *testing.T) {
	verification := &value_objects.Verification{
		Expiry: time.Now().Add(10 * time.Minute),
	}
	assert.True(t, verification.IsValid())
}

// TestIsValid_False は有効期限が現在時刻よりも前の場合、IsValid メソッドが false を返す。
func TestIsValid_False(t *testing.T) {
	verification := &value_objects.Verification{
		Expiry: time.Now().Add(-10 * time.Minute),
	}
	assert.False(t, verification.IsValid())
}
