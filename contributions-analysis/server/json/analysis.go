package json

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wnn-dev/contributions-analysis/handlers/analysis"
	"github.com/wnn-dev/contributions-analysis/objects"
	"github.com/wnn-dev/contributions-analysis/responder"
)

type AnalysisJsonServer struct {
	analysis *analysis.Handler
}

// NewAnalysisServer is an AnalysisResult JSON server constructor
func NewAnalysisServer(h *analysis.Handler) *AnalysisJsonServer {
	return &AnalysisJsonServer{analysis: h}
}

// Submit receives a contribution for analysis and creates a pending result
func (s *AnalysisJsonServer) Submit() gin.HandlerFunc {
	return func(c *gin.Context) {
		var vm *objects.AnalysisResultVM

		if err := c.ShouldBindJSON(&vm); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result, err := s.analysis.Submit(c, vm)
		if err != nil {
			responder.JsonResponse(c, false, err.Error(), nil)
			return
		}

		responder.JsonResponse(c, true, responder.ResourceCreated, result)
	}
}

// GetAnalysisResult returns a single analysis result by id
func (s *AnalysisJsonServer) GetAnalysisResult() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Query("id")
		result, err := s.analysis.AnalysisResult(c, id)
		if err != nil {
			responder.JsonResponse(c, false, err.Error(), nil)
			return
		}

		responder.JsonResponse(c, true, responder.ResourceFetched, result)
	}
}

// GetAnalysisResults returns all analysis results
func (s *AnalysisJsonServer) GetAnalysisResults() gin.HandlerFunc {
	return func(c *gin.Context) {
		results, err := s.analysis.AnalysisResults(c)
		if err != nil {
			responder.JsonResponse(c, false, err.Error(), nil)
			return
		}

		responder.JsonResponse(c, true, responder.ResourcesFetched, results)
	}
}

// GetByContributor returns all analysis results for a given contributor
func (s *AnalysisJsonServer) GetByContributor() gin.HandlerFunc {
	return func(c *gin.Context) {
		contributorID := c.Query("contributor_id")
		results, err := s.analysis.ByContributor(c, contributorID)
		if err != nil {
			responder.JsonResponse(c, false, err.Error(), nil)
			return
		}

		responder.JsonResponse(c, true, responder.ResourcesFetched, results)
	}
}

// GetByContribution returns all analysis results for a given contribution
func (s *AnalysisJsonServer) GetByContribution() gin.HandlerFunc {
	return func(c *gin.Context) {
		contributionID := c.Query("contribution_id")
		results, err := s.analysis.ByContribution(c, contributionID)
		if err != nil {
			responder.JsonResponse(c, false, err.Error(), nil)
			return
		}

		responder.JsonResponse(c, true, responder.ResourcesFetched, results)
	}
}

// UpdateStatus updates the status, score, and feedback of an analysis result
func (s *AnalysisJsonServer) UpdateStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Query("id")

		var vm *objects.AnalysisStatusUpdateVM
		if err := c.ShouldBindJSON(&vm); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := s.analysis.UpdateStatus(c, id, vm)
		if err != nil {
			responder.JsonResponse(c, false, err.Error(), nil)
			return
		}

		responder.JsonResponse(c, true, responder.ResourceUpdated, nil)
	}
}
