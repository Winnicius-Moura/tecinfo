package json

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wnn-dev/contributions-analysis/handlers/htmlcss"
	"github.com/wnn-dev/contributions-analysis/objects"
	"github.com/wnn-dev/contributions-analysis/responder"
)

type HtmlCssJsonServer struct {
	htmlcss *htmlcss.Handler
}

// NewHtmlCssServer is an HTML/CSS test JSON server constructor
func NewHtmlCssServer(h *htmlcss.Handler) *HtmlCssJsonServer {
	return &HtmlCssJsonServer{htmlcss: h}
}

// Submit accepts an HTML/CSS submission, runs the analyzer, and returns the report
func (s *HtmlCssJsonServer) Submit() gin.HandlerFunc {
	return func(c *gin.Context) {
		var vm *objects.HtmlCssSubmissionVM

		if err := c.ShouldBindJSON(&vm); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		report, err := s.htmlcss.Submit(c, vm)
		if err != nil {
			responder.JsonResponse(c, false, err.Error(), nil)
			return
		}

		responder.JsonResponse(c, true, responder.ResourceCreated, report)
	}
}

// GetSubmissionsByContributor returns all HTML/CSS submissions for a contributor
func (s *HtmlCssJsonServer) GetSubmissionsByContributor() gin.HandlerFunc {
	return func(c *gin.Context) {
		contributorID := c.Query("contributor_id")
		results, err := s.htmlcss.SubmissionsByContributor(c, contributorID)
		if err != nil {
			responder.JsonResponse(c, false, err.Error(), nil)
			return
		}

		responder.JsonResponse(c, true, responder.ResourcesFetched, results)
	}
}
