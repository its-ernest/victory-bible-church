package models

// ghost bug fixed: pointer-based fields for nullable columns
type Member struct {
    ID           string    `json:"id"`
    Phone        string    `json:"phone"`
    FirstName    *string    `json:"first_name"`
    LastName     *string    `json:"last_name"`
    Email        *string    `json:"email"`
    StatusID     int       `json:"status_id"`
    Status    string `json:"status"`
}