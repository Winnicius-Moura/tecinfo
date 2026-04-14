package responder

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserError is an error safe to deliver directly to the user
type UserError struct {
	code string
	err  string
}

func (e *UserError) Code() string {
	if e == nil {
		return ""
	}
	return e.code
}

func (e *UserError) Error() string {
	if e == nil {
		return ""
	}
	return e.err
}

func (e UserError) MarshalJSON() ([]byte, error) {
	j := struct {
		Code    string `json:"code,omitempty"`
		Message string `json:"message,omitempty"`
	}{
		Code:    e.code,
		Message: e.Error(),
	}
	return json.Marshal(&j)
}

var (
	ResourceCreated  = "Resource has been posted successfully"
	ResourceFetched  = "Resource is fetched successfully"
	ResourcesFetched = "Resources are fetched successfully"
	ResourceUpdated  = "Resource has been updated successfully"
	ResourceDeleted  = "Resource has been deleted successfully"
)

type ResponseBody struct {
	Code    int         `json:"status"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func JsonResponse(c *gin.Context, success bool, message string, data interface{}) {
	statusCode := http.StatusOK
	if !success {
		statusCode = http.StatusBadRequest
	}

	c.JSON(statusCode, ResponseBody{
		Code:    statusCode,
		Success: success,
		Message: message,
		Data:    data,
	})
}
