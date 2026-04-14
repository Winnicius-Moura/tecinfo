package objects

import (
	"context"
	"time"

	"github.com/wnn-dev/contributions-analysis/password"
)

type ContributorService interface {
	Create(ctx context.Context, data *Contributor) error
	Contributors(ctx context.Context) ([]*Contributor, error)
	Contributor(ctx context.Context, id string) (*Contributor, error)
	FindByEmail(ctx context.Context, email string) (*Contributor, error)
	Update(ctx context.Context, data *Contributor) error
	Delete(ctx context.Context, id string) error
}

type Contributor struct {
	ID            string         `gorm:"column:id;primary_key"                                      json:"id,omitempty"`
	Counter       int            `gorm:"column:counter"                                             json:"counter"`
	FullName      string         `gorm:"column:full_name"                                           json:"full_name,omitempty"`
	Email         string         `gorm:"column:email;unique_index;not null"                         json:"email"`
	Password      *password.Hash `gorm:"column:password;type:jsonb not null default '{}'::jsonb"    json:"password,omitempty"`
	Contributions []Contribution `gorm:"foreignkey:ContributorID"                                   json:"contributions,omitempty"`
	CreatedAt     time.Time      `gorm:"column:created_at"                                          json:"created_at,omitempty"`
	UpdatedAt     time.Time      `gorm:"column:updated_at"                                          json:"updated_at,omitempty"`
	DeletedAt     *time.Time     `gorm:"column:deleted_at"                                          json:"deleted_at,omitempty"`
}

type ContributorRegistrationVM struct {
	FullName string `json:"full_name,omitempty"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ContributorLoginVM struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
