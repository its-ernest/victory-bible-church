package models

import "time"

type Sermon struct {
	ID         string    `json:"id"`
	Title      string    `json:"title"`
	Speaker    string    `json:"speaker"`
	Series     string    `json:"series,omitempty"`
	VideoURL   string    `json:"video_url,omitempty"`
	AudioURL   string    `json:"audio_url,omitempty"`
	Summary    string    `json:"summary,omitempty"`
	Tags       []string  `json:"tags"`
	PreachedAt time.Time `json:"preached_at"`
}
