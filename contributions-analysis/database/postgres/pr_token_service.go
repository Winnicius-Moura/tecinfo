package postgres

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/wnn-dev/contributions-analysis/objects"
)

type PrTokenService struct {
	client *Client
}

func NewPrTokenService(ctx context.Context, client *Client) *PrTokenService {
	return &PrTokenService{client}
}

func (s *PrTokenService) Create(ctx context.Context, input *objects.PrToken) error {
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}
	input.ID = id.String()
	input.CreatedAt = time.Now()
	input.UpdatedAt = time.Now()
	input.Used = false

	return s.client.db.Create(input).Error
}

func (s *PrTokenService) Token(ctx context.Context, tokenStr string) (*objects.PrToken, error) {
	var token objects.PrToken
	err := s.client.db.Where("token = ?", tokenStr).First(&token).Error
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (s *PrTokenService) MarkAsUsed(ctx context.Context, id string) error {
	return s.client.db.Model(&objects.PrToken{}).Where("id = ?", id).Update("used", true).Error
}

func (s *PrTokenService) ByContributor(ctx context.Context, contributorID string) ([]*objects.PrToken, error) {
	var tokens []*objects.PrToken
	return tokens, s.client.db.Where("contributor_id = ?", contributorID).Order("created_at desc").Find(&tokens).Error
}
