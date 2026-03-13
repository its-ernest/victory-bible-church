package service

import (
	"context"
	"errors"
	"time"
	"log"
	"fmt"

	"church-backend/internal/utils"
	"github.com/its-ernest/echox/store"
	"church-backend/internal/repository"
)

type AuthService struct {
	store store.Store
	repo  *repository.MemberRepository
}

func NewAuthService(s store.Store, r *repository.MemberRepository) *AuthService {
	return &AuthService{store: s, repo: r}
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

	if err := s.repo.EnsureMember(ctx, phone); err != nil {
        log.Printf("[DB ERROR] Failed to ensure member %s: %v", phone, err)
        return fmt.Errorf("could not sync user record: %v", err)
    }
	return nil
}