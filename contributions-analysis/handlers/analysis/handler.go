package analysis

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/wnn-dev/contributions-analysis/objects"
)

type Handler struct {
	analysisResultService objects.AnalysisResultService
}

// NewHandler is the AnalysisResult handler constructor
func NewHandler(analysisResultService objects.AnalysisResultService) *Handler {
	return &Handler{analysisResultService: analysisResultService}
}

func (h *Handler) Submit(ctx *gin.Context, input *objects.AnalysisResultVM) (*objects.AnalysisResult, error) {
	result := &objects.AnalysisResult{
		ContributorID:  input.ContributorID,
		ContributionID: input.ContributionID,
	}

	err := h.analysisResultService.Create(ctx, result)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error submitting contribution for analysis"))
	}

	return result, nil
}

func (h *Handler) AnalysisResult(ctx *gin.Context, id string) (*objects.AnalysisResult, error) {
	result, err := h.analysisResultService.AnalysisResult(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error fetching analysis result"))
	}
	return result, nil
}

func (h *Handler) AnalysisResults(ctx *gin.Context) ([]*objects.AnalysisResult, error) {
	results, err := h.analysisResultService.AnalysisResults(ctx)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error fetching analysis results"))
	}
	return results, nil
}

func (h *Handler) ByContributor(ctx *gin.Context, contributorID string) ([]*objects.AnalysisResult, error) {
	results, err := h.analysisResultService.ByContributor(ctx, contributorID)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error fetching analysis results by contributor"))
	}
	return results, nil
}

func (h *Handler) ByContribution(ctx *gin.Context, contributionID string) ([]*objects.AnalysisResult, error) {
	results, err := h.analysisResultService.ByContribution(ctx, contributionID)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error fetching analysis results by contribution"))
	}
	return results, nil
}

func (h *Handler) UpdateStatus(ctx *gin.Context, id string, input *objects.AnalysisStatusUpdateVM) error {
	err := h.analysisResultService.UpdateStatus(ctx, id, input.Status, input.Score, input.Feedback)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("error updating analysis result status"))
	}
	return nil
}
