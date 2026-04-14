package postgres

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/wnn-dev/contributions-analysis/objects"
)

type ContributorService struct {
	client *Client
}

// NewContributorService is a Contributor service constructor
func NewContributorService(ctx context.Context, client *Client) *ContributorService {
	return &ContributorService{client}
}

func (s *ContributorService) Create(ctx context.Context, input *objects.Contributor) error {
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}
	input.ID = id.String()
	input.CreatedAt = time.Now()
	input.UpdatedAt = time.Now()

	return s.client.db.Create(input).Error
}

func (s *ContributorService) Contributors(ctx context.Context) ([]*objects.Contributor, error) {
	var contributors []*objects.Contributor
	return contributors, s.client.db.Order("created_at desc").Find(&contributors).Error
}

func (s *ContributorService) Contributor(ctx context.Context, id string) (*objects.Contributor, error) {
	var contributor objects.Contributor
	return &contributor, s.client.db.Set("gorm:auto_preload", true).Where("id = ?", id).Find(&contributor).Error
}

func (s *ContributorService) FindByEmail(ctx context.Context, email string) (*objects.Contributor, error) {
	var contributor objects.Contributor
	return &contributor, s.client.db.Set("gorm:auto_preload", true).Find(&contributor, "email = ?", email).Error
}

func (s *ContributorService) Update(ctx context.Context, input *objects.Contributor) error {
	input.UpdatedAt = time.Now()
	return s.client.db.Model(input).Where("id = ?", input.ID).
		Updates(map[string]interface{}{"full_name": input.FullName}).Error
}

func (s *ContributorService) Delete(ctx context.Context, id string) error {
	var contributor objects.Contributor
	return s.client.db.Where("id = ?", id).Delete(&contributor).Error
}
