package services

import (
	"InterviewPracticeBot-BE/internal/domain/entities"
	"InterviewPracticeBot-BE/internal/domain/repositories"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type EmailVerificationService struct {
	secretKey []byte
	repo      repositories.IUserPrivateRepository
}

var _ IEmailVerificationService = (*EmailVerificationService)(nil)

func NewEmailVerificationService(secretKey []byte, repo repositories.IUserPrivateRepository) *EmailVerificationService {
	return &EmailVerificationService{
		secretKey: secretKey,
		repo:      repo,
	}
}

func (s *EmailVerificationService) GenerateToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString(s.secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *EmailVerificationService) VerifyToken(tokenString string) (*entities.UserPrivate, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, jwt.ErrSignatureInvalid
		}
		return s.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := claims["userID"].(string)
		user, err := s.repo.FindByID(userID)
		if err != nil {
			return nil, err
		}
		return user, nil
	}

	return nil, jwt.ErrSignatureInvalid
}
