package postgres

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/wnn-dev/contributions-analysis/objects"
)

type ContributionService struct {
	client *Client
}

// NewContributionService is a Contribution service constructor
func NewContributionService(ctx context.Context, client *Client) *ContributionService {
	return &ContributionService{client}
}

func (s *ContributionService) Create(ctx context.Context, input objects.Contribution) error {
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}
	input.ID = id.String()
	input.CreatedAt = time.Now()
	input.UpdatedAt = time.Now()

	return s.client.db.Create(&input).Error
}

func (s *ContributionService) Contributions(ctx context.Context) ([]*objects.Contribution, error) {
	var contributions []*objects.Contribution
	return contributions, s.client.db.Order("created_at desc").Find(&contributions).Error
}

func (s *ContributionService) Contribution(ctx context.Context, id string) (*objects.Contribution, error) {
	var contribution objects.Contribution
	return &contribution, s.client.db.Where("id = ?", id).Find(&contribution).Error
}

func (s *ContributionService) Update(ctx context.Context, input *objects.Contribution) error {
	input.UpdatedAt = time.Now()
	return s.client.db.Model(input).Where("id = ?", input.ID).
		Updates(map[string]interface{}{
			"title":       input.Title,
			"description": input.Description,
			"kind":        input.Kind,
			"repository":  input.Repository,
		}).Error
}

func (s *ContributionService) Delete(ctx context.Context, id string) error {
	var contribution objects.Contribution
	return s.client.db.Where("id = ?", id).Delete(&contribution).Error
}
