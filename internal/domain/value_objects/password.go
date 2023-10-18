package value_objects

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	value string
}

func NewPassword(rawPassword string) (*Password, error) {
	if len(rawPassword) < 12 {
		return nil, errors.New("password must be at least 12 characters")
	}
	hashed, err := hashPassword(rawPassword)
	if err != nil {
		return nil, err
	}
	return &Password{value: hashed}, nil
}

func (p *Password) value() string {
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
