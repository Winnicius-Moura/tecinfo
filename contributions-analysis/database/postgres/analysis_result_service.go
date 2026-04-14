package postgres

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/wnn-dev/contributions-analysis/objects"
)

type AnalysisResultService struct {
	client *Client
}

// NewAnalysisResultService is an AnalysisResult service constructor
func NewAnalysisResultService(ctx context.Context, client *Client) *AnalysisResultService {
	return &AnalysisResultService{client}
}

func (s *AnalysisResultService) Create(ctx context.Context, input *objects.AnalysisResult) error {
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}
	input.ID = id.String()
	input.Status = objects.AnalysisStatusPending
	input.CreatedAt = time.Now()
	input.UpdatedAt = time.Now()

	return s.client.db.Create(input).Error
}

func (s *AnalysisResultService) AnalysisResults(ctx context.Context) ([]*objects.AnalysisResult, error) {
	var results []*objects.AnalysisResult
	return results, s.client.db.Order("created_at desc").Find(&results).Error
}

func (s *AnalysisResultService) AnalysisResult(ctx context.Context, id string) (*objects.AnalysisResult, error) {
	var result objects.AnalysisResult
	return &result, s.client.db.Set("gorm:auto_preload", true).Where("id = ?", id).First(&result).Error
}

func (s *AnalysisResultService) ByContributor(ctx context.Context, contributorID string) ([]*objects.AnalysisResult, error) {
	var results []*objects.AnalysisResult
	return results, s.client.db.Order("created_at desc").Where("contributor_id = ?", contributorID).Find(&results).Error
}

func (s *AnalysisResultService) ByContribution(ctx context.Context, contributionID string) ([]*objects.AnalysisResult, error) {
	var results []*objects.AnalysisResult
	return results, s.client.db.Order("created_at desc").Where("contribution_id = ?", contributionID).Find(&results).Error
}

func (s *AnalysisResultService) UpdateStatus(ctx context.Context, id string, status objects.AnalysisStatus, score float64, feedback string) error {
	return s.client.db.Model(&objects.AnalysisResult{}).Where("id = ?", id).
		Updates(map[string]interface{}{
			"status":     status,
			"score":      score,
			"feedback":   feedback,
			"updated_at": time.Now(),
		}).Error
}

func (s *AnalysisResultService) Delete(ctx context.Context, id string) error {
	var result objects.AnalysisResult
	return s.client.db.Where("id = ?", id).Delete(&result).Error
}
