package contribution

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/wnn-dev/contributions-analysis/objects"
)

type Handler struct {
	contributionService objects.ContributionService
}

// NewHandler is the Contribution handler constructor
func NewHandler(contributionService objects.ContributionService) *Handler {
	return &Handler{contributionService: contributionService}
}

func (h *Handler) Create(ctx *gin.Context, input *objects.ContributionVM) (*objects.Contribution, error) {
	contribution := objects.Contribution{
		Title:         input.Title,
		Description:   input.Description,
		Kind:          input.Kind,
		Repository:    input.Repository,
		ContributorID: input.ContributorID,
	}

	err := h.contributionService.Create(ctx, contribution)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error creating contribution"))
	}

	return &contribution, nil
}

func (h *Handler) Contribution(ctx *gin.Context, id string) (*objects.Contribution, error) {
	contribution, err := h.contributionService.Contribution(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error fetching contribution"))
	}
	return contribution, nil
}

func (h *Handler) Contributions(ctx *gin.Context) ([]*objects.Contribution, error) {
	contributions, err := h.contributionService.Contributions(ctx)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error fetching contributions"))
	}
	return contributions, nil
}
