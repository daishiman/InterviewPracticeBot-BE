package services_test

import (
	"InterviewPracticeBot-BE/internal/domain/entities"
	"InterviewPracticeBot-BE/internal/domain/repositories"
	"InterviewPracticeBot-BE/internal/domain/value_objects"
	"InterviewPracticeBot-BE/internal/services"
	"fmt"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

type mockUserRepo struct {
	user *entities.UserPrivate
}

func (m *mockUserRepo) FindByID(id string) (*entities.UserPrivate, error) {
	if m.user != nil && m.user.ID == id {
		return m.user, nil
	}
	return nil, repositories.ErrUserNotFound
}

func (m *mockUserRepo) FindByEmail(email string) (*entities.UserPrivate, error) {
	emailVO, err := value_objects.NewEmail(email)
	if err != nil {
		return nil, err
	}
	if m.user != nil && m.user.Email.Value() == emailVO.Value() {
		return m.user, nil
	}
	return nil, repositories.ErrUserNotFound
}

func (m *mockUserRepo) Save(user *entities.UserPrivate) error {
	m.user = user
	return nil
}

func (m *mockUserRepo) Delete(id string) error {
	if m.user != nil && m.user.ID == id {
		m.user = nil
		return nil
	}
	return repositories.ErrUserNotFound
}

func TestGenerateToken(t *testing.T) {
	secretKey := []byte("testSecretKey")
	userID := "testUserID"
	service := services.NewEmailVerificationService(secretKey, &mockUserRepo{})

	token, err := service.GenerateToken(userID)
	fmt.Println("service:", service, "token:", token, "err:", err)
	assert.NoError(t, err)

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secretKey), nil
	})
	assert.NoError(t, err)

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	assert.True(t, ok)
	assert.Equal(t, userID, claims["userID"])
	assert.LessOrEqual(t, time.Now().Unix(), int64(claims["exp"].(float64)))
}

func TestVerifyToken(t *testing.T) {
	secretKey := []byte("testSecretKey")
	userID := "testUserID"
	user := &entities.UserPrivate{ID: userID}
	service := services.NewEmailVerificationService(secretKey, &mockUserRepo{user: user})

	token, err := service.GenerateToken(userID)
	assert.NoError(t, err)

	verifiedUser, err := service.VerifyToken(token)
	assert.NoError(t, err)
	assert.Equal(t, user, verifiedUser)
}

func TestVerifyTokenWithInvalidToken(t *testing.T) {
	secretKey := []byte("testSecretKey")
	service := services.NewEmailVerificationService(secretKey, &mockUserRepo{})

	_, err := service.VerifyToken("invalidToken")
	assert.Error(t, err)
}

func TestVerifyTokenWithNonexistentUser(t *testing.T) {
	secretKey := []byte("testSecretKey")
	userID := "nonexistentUserID"
	service := services.NewEmailVerificationService(secretKey, &mockUserRepo{})

	token, err := service.GenerateToken(userID)
	assert.NoError(t, err)

	_, err = service.VerifyToken(token)
	assert.Error(t, err)
}
