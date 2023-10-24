package value_objects

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"time"
)

type Verification struct {
	Token  string
	Expiry time.Time
}

func NewVerification() (*Verification, error) {
	token, err := generateToken()
	if err != nil {
		return nil, errors.New("failed to generate token")
	}
	return &Verification{
		Token:  token,
		Expiry: time.Now().Add(1 * time.Hour),
	}, nil
}

func (v *Verification) IsValid() bool {
	return time.Now().Before(v.Expiry)
}

// トークン生成を行う関数
func generateToken() (string, error) {
	token := make([]byte, 32)
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(token), nil
}
