package contributor

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/wnn-dev/contributions-analysis/objects"
	"github.com/wnn-dev/contributions-analysis/password"
)

type Handler struct {
	contributorService  objects.ContributorService
	contributionService objects.ContributionService
}

// NewHandler is the Contributor handler constructor
func NewHandler(
	contributorService objects.ContributorService,
	contributionService objects.ContributionService,
) *Handler {
	return &Handler{
		contributorService:  contributorService,
		contributionService: contributionService,
	}
}

func (h *Handler) SignUp(ctx *gin.Context, input *objects.ContributorRegistrationVM) (*objects.Contributor, error) {
	pHashed, err := password.NewHashedPassword(input.Password)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error hashing password"))
	}

	contributor := &objects.Contributor{
		FullName: input.FullName,
		Email:    input.Email,
		Password: pHashed,
	}

	err = h.contributorService.Create(ctx, contributor)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error creating contributor"))
	}

	return contributor, nil
}

func (h *Handler) Login(ctx *gin.Context, input *objects.ContributorLoginVM) (*objects.Contributor, error) {
	contributor, err := h.contributorService.FindByEmail(ctx, input.Email)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error fetching contributor by email"))
	}

	if !contributor.Password.IsEqualTo(input.Password) {
		return nil, errors.New("invalid credentials")
	}

	return contributor, nil
}

func (h *Handler) Contributor(ctx *gin.Context, id string) (*objects.Contributor, error) {
	contributor, err := h.contributorService.Contributor(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error fetching contributor"))
	}
	return contributor, nil
}

func (h *Handler) Contributors(ctx *gin.Context) ([]*objects.Contributor, error) {
	contributors, err := h.contributorService.Contributors(ctx)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error fetching contributors"))
	}
	return contributors, nil
}
