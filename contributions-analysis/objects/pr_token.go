package objects

import (
	"context"
	"time"
)

type PrTokenService interface {
	Create(ctx context.Context, data *PrToken) error
	Token(ctx context.Context, tokenStr string) (*PrToken, error)
	MarkAsUsed(ctx context.Context, id string) error
	ByContributor(ctx context.Context, contributorID string) ([]*PrToken, error)
}

type PrToken struct {
	ID            string     `gorm:"column:id;primary_key"                  json:"id,omitempty"`
	ContributorID string     `gorm:"column:contributor_id;not null"         json:"contributor_id"`
	Token         string     `gorm:"column:token;unique_index;not null"     json:"token"`
	Used          bool       `gorm:"column:used;default:false"              json:"used"`
	ExpiresAt     time.Time  `gorm:"column:expires_at"                      json:"expires_at"`
	CreatedAt     time.Time  `gorm:"column:created_at"                      json:"created_at,omitempty"`
	UpdatedAt     time.Time  `gorm:"column:updated_at"                      json:"updated_at,omitempty"`
	DeletedAt     *time.Time `gorm:"column:deleted_at"                      json:"deleted_at,omitempty"`
}

type PrTokenValidationVM struct {
	Token string `json:"token"`
}
