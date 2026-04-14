package postgres

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/wnn-dev/contributions-analysis/objects"
)

type HtmlCssSubmissionService struct {
	client *Client
}

// NewHtmlCssSubmissionService is an HtmlCssSubmission service constructor
func NewHtmlCssSubmissionService(ctx context.Context, client *Client) *HtmlCssSubmissionService {
	return &HtmlCssSubmissionService{client}
}

func (s *HtmlCssSubmissionService) Create(ctx context.Context, input *objects.HtmlCssSubmission) error {
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}
	input.ID = id.String()
	input.CreatedAt = time.Now()
	input.UpdatedAt = time.Now()
	return s.client.db.Create(input).Error
}

func (s *HtmlCssSubmissionService) Submission(ctx context.Context, id string) (*objects.HtmlCssSubmission, error) {
	var result objects.HtmlCssSubmission
	return &result, s.client.db.Where("id = ?", id).First(&result).Error
}

func (s *HtmlCssSubmissionService) ByContributor(ctx context.Context, contributorID string) ([]*objects.HtmlCssSubmission, error) {
	var results []*objects.HtmlCssSubmission
	return results, s.client.db.Where("contributor_id = ?", contributorID).Order("created_at desc").Find(&results).Error
}
