package repository

import (
	"context"

	"church-backend/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type MinistryRepository struct {
	Pool *pgxpool.Pool
}

func (r *MinistryRepository) GetAllMinistries(ctx context.Context) ([]models.Ministry, error) {
	rows, err := r.Pool.Query(ctx, "SELECT id, name, description FROM church.ministries ORDER BY name ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.Ministry
	for rows.Next() {
		var m models.Ministry
		if err := rows.Scan(&m.ID, &m.Name, &m.Description); err != nil {
			return nil, err
		}
		list = append(list, m)
	}
	return list, nil
}

func (r *MinistryRepository) JoinMinistry(ctx context.Context, memberPhone, ministryID string) error {
	query := `
        INSERT INTO church.member_ministries (member_id, ministry_id)
        SELECT id, $2 
        FROM church.members 
        WHERE phone = $1
        ON CONFLICT (member_id, ministry_id) DO NOTHING
    `

	result, err := r.Pool.Exec(ctx, query, memberPhone, ministryID)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		// nothing now
	}
	return nil
}
