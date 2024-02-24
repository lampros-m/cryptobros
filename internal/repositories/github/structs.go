package github

import "time"

type Repository struct {
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updated_at"`
	Stars     int32     `json:"stargazers_count"`
}

type Repositories []Repository
