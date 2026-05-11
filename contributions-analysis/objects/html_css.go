package objects

import (
	"context"
	"time"
)

// HtmlCssSubmissionService interface for persistence operations
type HtmlCssSubmissionService interface {
	Create(ctx context.Context, data *HtmlCssSubmission) error
	Submission(ctx context.Context, id string) (*HtmlCssSubmission, error)
	ByContributor(ctx context.Context, contributorID string) ([]*HtmlCssSubmission, error)
	GallerySubmissions(ctx context.Context, limit int) ([]*GalleryCard, error)
}

// GalleryCard represents a summary of an approved submission for the gallery
type GalleryCard struct {
	ContributorID string    `json:"contributor_id"`
	HtmlContent   string    `json:"html_content"`
	ApprovedAt    time.Time `json:"approved_at"`
	Percentage    float64   `json:"percentage"`
}

// HtmlCssSubmission stores the raw HTML/CSS code submitted by a contributor
type HtmlCssSubmission struct {
	ID               string     `gorm:"column:id;primary_key"              json:"id,omitempty"`
	ContributorID    string     `gorm:"column:contributor_id;not null"     json:"contributor_id"`
	AnalysisResultID string     `gorm:"column:analysis_result_id"          json:"analysis_result_id,omitempty"`
	HtmlContent      string     `gorm:"column:html_content;type:text"      json:"html_content"`
	CreatedAt        time.Time  `gorm:"column:created_at"                  json:"created_at,omitempty"`
	UpdatedAt        time.Time  `gorm:"column:updated_at"                  json:"updated_at,omitempty"`
	DeletedAt        *time.Time `gorm:"column:deleted_at"                  json:"deleted_at,omitempty"`
}

// HtmlCssSubmissionVM is the input payload from the student
type HtmlCssSubmissionVM struct {
	ContributorID string `json:"contributor_id"`
	HtmlContent   string `json:"html_content"`
}

// CheckResult represents the outcome of a single analysis rule
type CheckResult struct {
	Rule      string `json:"rule"`
	Passed    bool   `json:"passed"`
	Points    int    `json:"points"`
	MaxPoints int    `json:"max_points"`
	Expected  string `json:"expected,omitempty"`
	Actual    string `json:"actual,omitempty"`
	Diff      string `json:"diff,omitempty"`
}

// HtmlCssAnalysisReport is the full analysis result returned to the contributor
type HtmlCssAnalysisReport struct {
	Score        float64       `json:"score"`
	MaxScore     float64       `json:"max_score"`
	Percentage   float64       `json:"percentage"`
	Approved     bool          `json:"approved"`
	PRToken      string        `json:"pr_token,omitempty"`
	PassedChecks []CheckResult `json:"passed_checks"`
	FailedChecks []CheckResult `json:"failed_checks"`
}
