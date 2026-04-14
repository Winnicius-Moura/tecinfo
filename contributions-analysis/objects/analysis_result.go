package objects

import (
	"context"
	"time"
)

// AnalysisStatus represents the possible states of a contribution analysis
type AnalysisStatus string

const (
	AnalysisStatusPending  AnalysisStatus = "pending"
	AnalysisStatusRunning  AnalysisStatus = "running"
	AnalysisStatusApproved AnalysisStatus = "approved"
	AnalysisStatusRejected AnalysisStatus = "rejected"
)

type AnalysisResultService interface {
	Create(ctx context.Context, data *AnalysisResult) error
	AnalysisResults(ctx context.Context) ([]*AnalysisResult, error)
	AnalysisResult(ctx context.Context, id string) (*AnalysisResult, error)
	ByContributor(ctx context.Context, contributorID string) ([]*AnalysisResult, error)
	ByContribution(ctx context.Context, contributionID string) ([]*AnalysisResult, error)
	UpdateStatus(ctx context.Context, id string, status AnalysisStatus, score float64, feedback string) error
	Delete(ctx context.Context, id string) error
}

// AnalysisResult stores the result of analyzing a contribution
type AnalysisResult struct {
	ID             string         `gorm:"column:id;primary_key"             json:"id,omitempty"`
	ContributorID  string         `gorm:"column:contributor_id;not null"    json:"contributor_id"`
	ContributionID string         `gorm:"column:contribution_id;not null"   json:"contribution_id"`
	Status         AnalysisStatus `gorm:"column:status;not null"            json:"status"`
	Score          float64        `gorm:"column:score"                      json:"score"`
	Feedback       string         `gorm:"column:feedback;type:text"         json:"feedback,omitempty"`
	CreatedAt      time.Time      `gorm:"column:created_at"                 json:"created_at,omitempty"`
	UpdatedAt      time.Time      `gorm:"column:updated_at"                 json:"updated_at,omitempty"`
	DeletedAt      *time.Time     `gorm:"column:deleted_at"                 json:"deleted_at,omitempty"`

	// Preloaded associations
	Contributor  *Contributor  `gorm:"foreignkey:ContributorID"  json:"contributor,omitempty"`
	Contribution *Contribution `gorm:"foreignkey:ContributionID" json:"contribution,omitempty"`
}

type AnalysisResultVM struct {
	ContributorID  string `json:"contributor_id"`
	ContributionID string `json:"contribution_id"`
}

type AnalysisStatusUpdateVM struct {
	Status   AnalysisStatus `json:"status"`
	Score    float64        `json:"score"`
	Feedback string         `json:"feedback"`
}
