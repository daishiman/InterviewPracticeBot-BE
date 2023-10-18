// Package repositories provides interfaces for database access.
package repositories

import (
	"InterviewPracticeBot-BE/internal/domain/entities"
	"errors"
)

type InMemoryUserRepository struct {
	users map[string]*entities.UserPrivate
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*entities.UserPrivate),
	}
}

func (repo *InMemoryUserRepository) FindByID(id string) (*entities.UserPrivate, error) {
	user, exists := repo.users[id]
	if !exists {
		return nil, ErrUserNotFound
	}
	return user, nil
}

func (repo *InMemoryUserRepository) FindByEmail(email string) (*entities.UserPrivate, error) {
	for _, user := range repo.users {
		if user.Email.Value() == email {
			return user, nil
		}
	}
	return nil, ErrUserNotFound
}

func (repo *InMemoryUserRepository) Save(user *entities.UserPrivate) error {
	if _, exists := repo.users[user.ID]; exists {
		return errors.New("user already exists")
	}
	repo.users[user.ID] = user
	return nil
}

func (repo *InMemoryUserRepository) Delete(id string) error {
	if _, exists := repo.users[id]; !exists {
		return ErrUserNotFound
	}
	delete(repo.users, id)
	return nil
}
