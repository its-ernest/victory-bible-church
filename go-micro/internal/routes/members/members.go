package members

import (
	"context"
	"strings"
	"errors"

	"church-backend/internal/repository"
	"church-backend/internal/models"
)

type Service struct {
	repo *repository.MemberRepository
}

func NewService(r *repository.MemberRepository) *Service {
	return &Service{repo: r}
}

func (s *Service) GetProfile(ctx context.Context, phone string) (*models.Member, error) {
	return s.repo.FindByPhone(ctx, phone)
}

func (s *Service) UpdateMember(ctx context.Context, phone string, firstName, lastName, email string) error {
    firstName = strings.TrimSpace(firstName)
    lastName = strings.TrimSpace(lastName)
    email = strings.ToLower(strings.TrimSpace(email))

    // required fields
    if firstName == "" || lastName == "" {
        return errors.New("first name and last name are required")
    }

    // optional email validation
    var emailPtr *string
    if email != "" {
        if !strings.Contains(email, "@") {
            return errors.New("invalid email format")
        }
        emailPtr = &email
    } else {
        emailPtr = nil
    }

    m := &models.Member{
        Phone:     phone,
        FirstName: &firstName,
        LastName:  &lastName,
        Email:     emailPtr,
    }
    return s.repo.UpdateProfile(ctx, m)
}

func (s *Service) Search(ctx context.Context, query string, page, pageSize int) ([]models.Member, error) {
    if pageSize <= 0 { pageSize = 10 }
    if page <= 0 { page = 1 }
    
    offset := (page - 1) * pageSize
    return s.repo.SearchMembers(ctx, query, pageSize, offset)
}