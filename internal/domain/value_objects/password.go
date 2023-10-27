package value_objects

import (
	"errors"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

var (
	uppercaseRegex = regexp.MustCompile(`[A-Z]`)
	lowercaseRegex = regexp.MustCompile(`[a-z]`)
	digitRegex     = regexp.MustCompile(`[0-9]`)
	specialRegex   = regexp.MustCompile(`[!@#\$%\^&\*]`)
)

type Password struct {
	value string
}

func NewPassword(rawPassword string) (*Password, error) {
	if !isValidPassword(rawPassword) {
		return nil, errors.New("password must be at least 12 characters long, contain an uppercase letter, a lowercase letter, a digit, and a special character")
	}
	hashed, err := hashPassword(rawPassword)
	if err != nil {
		return nil, err
	}
	return &Password{value: hashed}, nil
}

func (p *Password) Compare(rawPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(p.value), []byte(rawPassword))
	return err == nil, err
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

func isValidPassword(password string) bool {
	hasUppercase := uppercaseRegex.MatchString(password)
	hasLowercase := lowercaseRegex.MatchString(password)
	hasDigit := digitRegex.MatchString(password)
	hasSpecial := specialRegex.MatchString(password)
	hasMinLength := len(password) >= 12

	return hasUppercase && hasLowercase && hasDigit && hasSpecial && hasMinLength
}
