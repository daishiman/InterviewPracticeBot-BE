package value_objects

import (
	"errors"
	"regexp"
	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	value string
}

func NewPassword(rawPassword string) (*Password, error) {
	if !isValidPassword(rawPassword) {
    return nil, errors.New("password must meet the complexity requirements")
	}
	hashed, err := hashPassword(rawPassword)
	if err != nil {
		return nil, err
	}
	return &Password{value: hashed}, nil
}

func (p *Password) Value() string {
	return p.value
}

func hashPassword(password string) (string, error) {
	const cost = 14
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func ComparePassword(hashedPassword, rawPassword string) (bool, error) {
  err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(rawPassword))
  return err == nil, err
}

func isValidPassword(password string) bool {
  // 大文字、小文字、数字、特殊文字をそれぞれ1回以上含むかどうかをチェック
  hasUppercase := regexp.MustCompile(`[A-Z]`).MatchString(password)
  hasLowercase := regexp.MustCompile(`[a-z]`).MatchString(password)
  hasDigit := regexp.MustCompile(`[0-9]`).MatchString(password)
  hasSpecial := regexp.MustCompile(`[!@#\$%\^&\*]`).MatchString(password)
  hasMinLength := len(password) >= 12

  return hasUppercase && hasLowercase && hasDigit && hasSpecial && hasMinLength
}


