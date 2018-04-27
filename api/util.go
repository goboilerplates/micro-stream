package api

import "github.com/gin-gonic/gin"

// HandleStreamResponse handles the Stream API response.
func HandleStreamResponse(context *gin.Context, body interface{}, err error, messageName string) bool {
	if err != nil {
		return false
	}
	context.SSEvent(messageName, body)
	return true
}
