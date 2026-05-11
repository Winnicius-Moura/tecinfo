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

// ValidatePRToken validates a PR token
func (s *HtmlCssJsonServer) ValidatePRToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var vm *objects.PrTokenValidationVM

		if err := c.ShouldBindJSON(&vm); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := s.htmlcss.ValidatePRToken(c, vm)
		if err != nil {
			responder.JsonResponse(c, false, err.Error(), nil)
			return
		}

		responder.JsonResponse(c, true, "Token validado com sucesso", nil)
	}
}

// GetPrTokensByContributor returns all PR tokens for a contributor
func (s *HtmlCssJsonServer) GetPrTokensByContributor() gin.HandlerFunc {
	return func(c *gin.Context) {
		contributorID := c.Query("contributor_id")
		results, err := s.htmlcss.PrTokensByContributor(c, contributorID)
		if err != nil {
			responder.JsonResponse(c, false, err.Error(), nil)
			return
		}

		responder.JsonResponse(c, true, responder.ResourcesFetched, results)
	}
}

// GetGallery returns the latest approved submissions for the gallery
func (s *HtmlCssJsonServer) GetGallery() gin.HandlerFunc {
	return func(c *gin.Context) {
		results, err := s.htmlcss.GallerySubmissions(c, 50)
		if err != nil {
			responder.JsonResponse(c, false, err.Error(), nil)
			return
		}

		responder.JsonResponse(c, true, responder.ResourcesFetched, results)
	}
}
