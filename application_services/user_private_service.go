package application_services

import (
	"InterviewPracticeBot-BE/internal/domain/aggregates"
	"InterviewPracticeBot-BE/internal/domain/entities"
	"InterviewPracticeBot-BE/internal/domain/repositories"
	"InterviewPracticeBot-BE/internal/domain/value_objects"
	"InterviewPracticeBot-BE/internal/services"
	"fmt"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

type UserPrivateService struct {
	userAggregate aggregates.IUserAggregate
	userRepo      repositories.IUserPrivateRepository
	emailService  services.IEmailVerificationService
}

func NewUserPrivateService(userAggregate aggregates.IUserAggregate, userRepo repositories.IUserPrivateRepository, emailService services.IEmailVerificationService) *UserPrivateService {
	return &UserPrivateService{
		userAggregate: userAggregate,
		userRepo:      userRepo,
		emailService:  emailService,
	}
}

func (s *UserPrivateService) RegisterUser(email value_objects.Email, password value_objects.Password) (entities.IUserPrivate, error) {
	fmt.Println("test")
	user, err := s.userAggregate.CreateUser(email.Value(), password.Value())
	if err != nil {
		return nil, err
	}

	userPrivate, ok := user.(*entities.UserPrivate)
	if !ok {
		return nil, fmt.Errorf("failed to assert user to *entities.UserPrivate")
	}

	err = s.userRepo.Save(userPrivate)
	if err != nil {
		return nil, err
	}

	token, err := s.emailService.GenerateToken(user.GetID())
	if err != nil {
		return nil, err
	}

	err = s.SendConfirmationEmail(user.GetID(), token)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserPrivateService) UpdateUserPassword(userID string, oldPassword, newPassword value_objects.Password) error {
	err := s.userAggregate.UpdatePassword(userID, oldPassword, newPassword)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserPrivateService) SendConfirmationEmail(userID, token string) error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("failed to load .env file: %w", err)
	}

	from := os.Getenv("FROM_EMAIL")
	to := userID // ここでは、userIDがメールアドレスと仮定しています
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	body := "Click on the link to verify your account: https://example.com/verify?token=" + token

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Account Verification\n\n" +
		body

	err = smtp.SendMail(smtpHost+":"+smtpPort, nil, from, []string{to}, []byte(msg))
	if err != nil {
		return fmt.Errorf("failed to send confirmation email: %w", err)
	}

	return nil
}
