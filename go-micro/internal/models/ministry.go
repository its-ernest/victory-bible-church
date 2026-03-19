package models

type Ministry struct {
    ID          string `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
}

type MemberMinistryRequest struct {
    MinistryID string `json:"ministry_id"`
}