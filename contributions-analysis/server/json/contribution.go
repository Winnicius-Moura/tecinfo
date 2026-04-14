package json

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wnn-dev/contributions-analysis/handlers/contribution"
	"github.com/wnn-dev/contributions-analysis/objects"
	"github.com/wnn-dev/contributions-analysis/responder"
)

type ContributionJsonServer struct {
	contribution *contribution.Handler
}

// NewContributionServer is a Contribution JSON server constructor
func NewContributionServer(h *contribution.Handler) *ContributionJsonServer {
	return &ContributionJsonServer{contribution: h}
}

func (s *ContributionJsonServer) CreateContribution() gin.HandlerFunc {
	return func(c *gin.Context) {
		var vm *objects.ContributionVM

		if err := c.ShouldBindJSON(&vm); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result, err := s.contribution.Create(c, vm)
		if err != nil {
			responder.JsonResponse(c, false, err.Error(), nil)
			return
		}

		responder.JsonResponse(c, true, responder.ResourceCreated, result)
	}
}

func (s *ContributionJsonServer) GetContribution() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Query("id")
		result, err := s.contribution.Contribution(c, id)
		if err != nil {
			responder.JsonResponse(c, false, err.Error(), nil)
			return
		}

		responder.JsonResponse(c, true, responder.ResourceFetched, result)
	}
}

func (s *ContributionJsonServer) GetContributions() gin.HandlerFunc {
	return func(c *gin.Context) {
		results, err := s.contribution.Contributions(c)
		if err != nil {
			responder.JsonResponse(c, false, err.Error(), nil)
			return
		}

		responder.JsonResponse(c, true, responder.ResourcesFetched, results)
	}
}
