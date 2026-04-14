package json

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wnn-dev/contributions-analysis/handlers/contributor"
	"github.com/wnn-dev/contributions-analysis/objects"
	"github.com/wnn-dev/contributions-analysis/responder"
)

type ContributorJsonServer struct {
	contributor *contributor.Handler
}

// NewContributorServer is a Contributor JSON server constructor
func NewContributorServer(h *contributor.Handler) *ContributorJsonServer {
	return &ContributorJsonServer{contributor: h}
}

func (s *ContributorJsonServer) SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var vm *objects.ContributorRegistrationVM

		if err := c.ShouldBindJSON(&vm); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result, err := s.contributor.SignUp(c, vm)
		if err != nil {
			responder.JsonResponse(c, false, err.Error(), nil)
			return
		}

		responder.JsonResponse(c, true, responder.ResourceCreated, result)
	}
}

func (s *ContributorJsonServer) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var vm *objects.ContributorLoginVM

		if err := c.ShouldBindJSON(&vm); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result, err := s.contributor.Login(c, vm)
		if err != nil {
			responder.JsonResponse(c, false, err.Error(), nil)
			return
		}

		responder.JsonResponse(c, true, responder.ResourceFetched, result)
	}
}

func (s *ContributorJsonServer) GetContributor() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Query("id")
		result, err := s.contributor.Contributor(c, id)
		if err != nil {
			responder.JsonResponse(c, false, err.Error(), nil)
			return
		}

		responder.JsonResponse(c, true, responder.ResourceFetched, result)
	}
}

func (s *ContributorJsonServer) GetContributors() gin.HandlerFunc {
	return func(c *gin.Context) {
		results, err := s.contributor.Contributors(c)
		if err != nil {
			responder.JsonResponse(c, false, err.Error(), nil)
			return
		}

		responder.JsonResponse(c, true, responder.ResourcesFetched, results)
	}
}
