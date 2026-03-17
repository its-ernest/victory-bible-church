package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"church-backend/internal/models"
)

type MemberRepository struct {
	Pool *pgxpool.Pool
}

// FindByPhone retrieves a full member profile including their status name
func (r *MemberRepository) FindByPhone(ctx context.Context, phone string) (*models.Member, error) {
	var m models.Member
	query := `
        SELECT m.id, m.phone, m.first_name, m.last_name, m.email, m.status_id, s.name as status
        FROM church.members m
        JOIN church.member_status s ON m.status_id = s.id
        WHERE m.phone = $1
    `
    err := r.Pool.QueryRow(ctx, query, phone).Scan(
        &m.ID, 
        &m.Phone, 
        &m.FirstName, 
        &m.LastName, 
        &m.Email, 
        &m.StatusID,
        &m.Status,
	)
	if err != nil {
		return nil, fmt.Errorf("member not found: %w", err)
	}
	return &m, nil
}

// UpdateProfile updates a member
func (r *MemberRepository) UpdateProfile(ctx context.Context, m *models.Member) error {
	query := `
        UPDATE church.members 
        SET first_name = $1, last_name = $2, email = $3, updated_at = NOW() 
        WHERE phone = $4
    `
	result, err := r.Pool.Exec(ctx, query, m.FirstName, m.LastName, m.Email, m.Phone)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("no member found with phone %s", m.Phone)
	}
	return nil
}

func (r *MemberRepository) EnsureMember(ctx context.Context, phone string) error {
    // default new signups to 'Visitor' (ID 2 in db seed)
    query := `
        INSERT INTO church.members (phone, status_id)
        VALUES ($1, (SELECT id FROM church.member_status WHERE name = 'Visitor'))
        ON CONFLICT (phone) DO NOTHING;
    `
    _, err := r.Pool.Exec(ctx, query, phone)
    return err
}

func (r *MemberRepository) SearchMembers(ctx context.Context, query string, limit, offset int) ([]models.Member, error) {
    var members []models.Member
    
    // search by name, email, or phone
    sql := `
        SELECT m.id, m.phone, m.first_name, m.last_name, m.email, m.status_id, s.name as status
        FROM church.members m
        JOIN church.member_status s ON m.status_id = s.id
        WHERE m.first_name ILIKE $1 
           OR m.last_name ILIKE $1 
           OR m.phone ILIKE $1 
           OR m.email ILIKE $1
        ORDER BY m.last_name ASC
        LIMIT $2 OFFSET $3
    `
    
    searchParam := "%" + query + "%"
    rows, err := r.Pool.Query(ctx, sql, searchParam, limit, offset)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var m models.Member
        err := rows.Scan(&m.ID, &m.Phone, &m.FirstName, &m.LastName, &m.Email, &m.StatusID, &m.Status)
        if err != nil {
            return nil, err
        }
        members = append(members, m)
    }
    
    return members, nil
}