package objects

import (
	"context"
	"time"
)

type ContributionService interface {
	Create(ctx context.Context, data Contribution) error
	Contributions(ctx context.Context) ([]*Contribution, error)
	Contribution(ctx context.Context, id string) (*Contribution, error)
	Update(ctx context.Context, data *Contribution) error
	Delete(ctx context.Context, id string) error
}

type Contribution struct {
	ID            string     `gorm:"column:id;primary_key"  json:"id,omitempty"`
	Counter       int        `gorm:"column:counter"         json:"counter"`
	Title         string     `gorm:"column:title"           json:"title,omitempty"`
	Description   string     `gorm:"column:description"     json:"description,omitempty"`
	Kind          string     `gorm:"column:kind"            json:"kind,omitempty"`
	Repository    string     `gorm:"column:repository"      json:"repository,omitempty"`
	ContributorID string     `gorm:"column:contributor_id"  json:"contributor_id"`
	CreatedAt     time.Time  `gorm:"column:created_at"      json:"created_at,omitempty"`
	UpdatedAt     time.Time  `gorm:"column:updated_at"      json:"updated_at,omitempty"`
	DeletedAt     *time.Time `gorm:"column:deleted_at"      json:"deleted_at,omitempty"`
}

type ContributionVM struct {
	Title         string `json:"title,omitempty"`
	Description   string `json:"description,omitempty"`
	Kind          string `json:"kind,omitempty"`
	Repository    string `json:"repository,omitempty"`
	ContributorID string `json:"contributor_id"`
}
