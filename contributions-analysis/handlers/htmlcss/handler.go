package htmlcss

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/wnn-dev/contributions-analysis/analysis/htmlcss"
	"github.com/wnn-dev/contributions-analysis/objects"
	"github.com/wnn-dev/contributions-analysis/rabbitmq"
)

type Handler struct {
	submissionService     objects.HtmlCssSubmissionService
	analysisResultService objects.AnalysisResultService
	prTokenService        objects.PrTokenService
	publisher             *rabbitmq.Publisher
}

// NewHandler is the HTML/CSS test handler constructor
func NewHandler(
	submissionService objects.HtmlCssSubmissionService,
	analysisResultService objects.AnalysisResultService,
	prTokenService objects.PrTokenService,
	publisher *rabbitmq.Publisher,
) *Handler {
	return &Handler{
		submissionService:     submissionService,
		analysisResultService: analysisResultService,
		prTokenService:        prTokenService,
		publisher:             publisher,
	}
}

func generatePRTokenString(contributorID string) string {
	year := time.Now().Year()
	idPrefix := contributorID
	if len(contributorID) > 4 {
		idPrefix = contributorID[:4]
	}
	bytes := make([]byte, 3)
	rand.Read(bytes)
	randomPart := hex.EncodeToString(bytes)
	return fmt.Sprintf("%s_tecInfo_%d_%s", idPrefix, year, randomPart)
}

// Submit runs the analysis engine, persists the submission and the result, and returns the report
func (h *Handler) Submit(ctx *gin.Context, vm *objects.HtmlCssSubmissionVM) (*objects.HtmlCssAnalysisReport, error) {
	// 1. Run analysis
	report, err := htmlcss.Analyze(vm.HtmlContent)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error analyzing HTML/CSS submission"))
	}

	// 2. Generate PR Token if approved
	if report.Approved {
		tokenStr := generatePRTokenString(vm.ContributorID)
		prToken := &objects.PrToken{
			ContributorID: vm.ContributorID,
			Token:         tokenStr,
			ExpiresAt:     time.Now().Add(24 * time.Hour),
		}
		if err := h.prTokenService.Create(ctx, prToken); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error saving PR token"))
		}
		report.PRToken = tokenStr
	}

	// 3. Serialize detailed report into feedback JSON
	feedbackJSON, err := json.Marshal(report)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error serializing analysis report"))
	}

	// 4. Determine approval status
	status := objects.AnalysisStatusRejected
	if report.Approved {
		status = objects.AnalysisStatusApproved
	}

	// 5. Persist analysis result
	analysisResult := &objects.AnalysisResult{
		ContributorID: vm.ContributorID,
		Status:        status,
		Score:         report.Percentage,
		Feedback:      string(feedbackJSON),
	}
	if err := h.analysisResultService.Create(ctx, analysisResult); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error saving analysis result"))
	}

	// 6. Persist the raw submission linked to the result
	submission := &objects.HtmlCssSubmission{
		ContributorID:    vm.ContributorID,
		AnalysisResultID: analysisResult.ID,
		HtmlContent:      vm.HtmlContent,
	}
	if err := h.submissionService.Create(ctx, submission); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error saving HTML/CSS submission"))
	}

	// 7. If approved, publish to RabbitMQ for real-time gallery
	if report.Approved && h.publisher != nil {
		err := h.publisher.PublishApprovedSubmission(ctx, submission, report)
		if err != nil {
			// Just log the error, don't fail the whole submit request
			fmt.Printf("Failed to publish to RabbitMQ: %v\n", err)
		}
	}

	return report, nil
}

// SubmissionsByContributor returns all HTML/CSS submissions for a given contributor
func (h *Handler) SubmissionsByContributor(ctx *gin.Context, contributorID string) ([]*objects.HtmlCssSubmission, error) {
	results, err := h.submissionService.ByContributor(ctx, contributorID)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error fetching submissions by contributor"))
	}
	return results, nil
}

// ValidatePRToken validates and consumes a PR token
func (h *Handler) ValidatePRToken(ctx *gin.Context, vm *objects.PrTokenValidationVM) error {
	token, err := h.prTokenService.Token(ctx, vm.Token)
	if err != nil {
		return errors.New("token inválido")
	}

	if token.Used {
		return errors.New("token já utilizado")
	}

	if time.Now().After(token.ExpiresAt) {
		return errors.New("token expirado")
	}

	return h.prTokenService.MarkAsUsed(ctx, token.ID)
}

// PrTokensByContributor returns all PR tokens for a given contributor
func (h *Handler) PrTokensByContributor(ctx *gin.Context, contributorID string) ([]*objects.PrToken, error) {
	tokens, err := h.prTokenService.ByContributor(ctx, contributorID)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error fetching pr tokens"))
	}
	return tokens, nil
}

// GallerySubmissions returns the latest approved submissions
func (h *Handler) GallerySubmissions(ctx *gin.Context, limit int) ([]*objects.GalleryCard, error) {
	results, err := h.submissionService.GallerySubmissions(ctx, limit)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error fetching gallery submissions"))
	}
	return results, nil
}
