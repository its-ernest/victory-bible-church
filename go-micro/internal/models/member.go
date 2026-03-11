type Member struct {
    ID           string    `json:"id"`
    Phone        string    `json:"phone"`
    FirstName    string    `json:"first_name"`
    LastName     string    `json:"last_name"`
    IsBaptized   bool      `json:"is_baptized"`
    StatusID     int       `json:"status_id"`
}