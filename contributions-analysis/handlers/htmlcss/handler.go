package htmlcss

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/wnn-dev/contributions-analysis/analysis/htmlcss"
	"github.com/wnn-dev/contributions-analysis/objects"
)

type Handler struct {
	submissionService     objects.HtmlCssSubmissionService
	analysisResultService objects.AnalysisResultService
}

// NewHandler is the HTML/CSS test handler constructor
func NewHandler(
	submissionService objects.HtmlCssSubmissionService,
	analysisResultService objects.AnalysisResultService,
) *Handler {
	return &Handler{
		submissionService:     submissionService,
		analysisResultService: analysisResultService,
	}
}

// Submit runs the analysis engine, persists the submission and the result, and returns the report
func (h *Handler) Submit(ctx *gin.Context, vm *objects.HtmlCssSubmissionVM) (*objects.HtmlCssAnalysisReport, error) {
	// 1. Run analysis
	report, err := htmlcss.Analyze(vm.HtmlContent)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error analyzing HTML/CSS submission"))
	}

	// 2. Serialize detailed report into feedback JSON
	feedbackJSON, err := json.Marshal(report)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error serializing analysis report"))
	}

	// 3. Determine approval status
	status := objects.AnalysisStatusRejected
	if report.Approved {
		status = objects.AnalysisStatusApproved
	}

	// 4. Persist analysis result
	analysisResult := &objects.AnalysisResult{
		ContributorID: vm.ContributorID,
		Status:        status,
		Score:         report.Percentage,
		Feedback:      string(feedbackJSON),
	}
	if err := h.analysisResultService.Create(ctx, analysisResult); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error saving analysis result"))
	}

	// 5. Persist the raw submission linked to the result
	submission := &objects.HtmlCssSubmission{
		ContributorID:    vm.ContributorID,
		AnalysisResultID: analysisResult.ID,
		HtmlContent:      vm.HtmlContent,
	}
	if err := h.submissionService.Create(ctx, submission); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error saving HTML/CSS submission"))
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
