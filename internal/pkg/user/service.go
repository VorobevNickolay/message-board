package user

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type store interface {
	CreateUser(ctx context.Context, name, password string) (User, error)
	FindUserByName(ctx context.Context, name string) (User, error)
	FindUserByID(ctx context.Context, id string) (User, error)
}

type Service struct {
	store store
}

func NewService(store store) *Service {
	return &Service{store: store}
}

func (s *Service) SignUp(ctx context.Context, name, password string) (User, error) {
	_, err := s.store.FindUserByName(ctx, name)
	if err == nil {
		return User{}, ErrUsedUsername
	}

	password = s.createHash(password)
	user, err := s.store.CreateUser(ctx, name, password)
	if err != nil {
		return User{}, fmt.Errorf("failed to signup: %w", err)
	}
	return user, nil
}

func (s *Service) Login(ctx context.Context, name, password string) (User, error) {
	u, err := s.store.FindUserByName(ctx, name)
	if err != nil {
		return User{}, err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return User{}, ErrUserNotFound
	}
	return u, nil
}

func (s *Service) createHash(str string) string {
	bytePassword := []byte(str)
	hashPassword, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	return string(hashPassword)
}

func (s *Service) FindUserByID(ctx context.Context, id string) (User, error) {
	u, err := s.store.FindUserByID(ctx, id)
	if err != nil {
		return User{}, err
	}
	return u, nil
}
