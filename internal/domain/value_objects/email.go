package value_objects

import (
	"errors"
	"regexp"
)

type Email struct {
	value string
}

func NewEmail(email string) (*Email, error) {
	if !isValidEmail(email) {
		return nil, errors.New("invalid email format")
	}
	return &Email{value: email}, nil
}

func (e *Email) Value() string {
	return e.value
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`)
	return re.MatchString(email)
}
