package services

import (
	"InterviewPracticeBot-BE/internal/domain/entities"
)

type IEmailVerificationService interface {
	GenerateToken(userId string) (string, error)
	VerifyToken(token string) (*entities.UserPrivate, error)
}
