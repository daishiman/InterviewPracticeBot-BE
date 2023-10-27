package application_services

import (
	"InterviewPracticeBot-BE/internal/domain/entities"
	"InterviewPracticeBot-BE/internal/domain/value_objects"
)

type IUserPrivateService interface {
	RegisterUser(email value_objects.Email, password value_objects.Password) (entities.IUserPrivate, error)
	UpdateUserPassword(userID string, oldPassword value_objects.Password, newPassword value_objects.Password) error
	SendConfirmationEmail(userID, token string) error
}
