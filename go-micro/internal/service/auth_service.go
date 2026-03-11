package service

import (
	"context"
	"errors"
	"time"

	"church-backend/internal/utils"
	"github.com/its-ernest/echox/store"
)

type AuthService struct {
	store store.Store
}

func NewAuthService(s store.Store) *AuthService {
	return &AuthService{store: s}
}

func (s *AuthService) RequestOTP(ctx context.Context, phone string) error {
	code := utils.GenerateOTP()

	// entry struct data wrapper
	entry := &store.Entry{
		Body: []byte(code),
	}

	return s.store.Save(ctx, "otp:"+phone, entry, 5*time.Minute)
}

func (s *AuthService) VerifyOTP(ctx context.Context, phone, code string) error {
	entry, err := s.store.Get(ctx, "otp:"+phone)
	if err != nil {
		return errors.New("otp expired or not found")
	}

	if string(entry.Body) != code {
		return errors.New("invalid otp: original: "+string(entry.Body))
	}

	_ = s.store.Delete(ctx, "otp:"+phone)
	return nil
}