package config

import (
	"github.com/gin-gonic/gin"
	"github.com/goboilerplates/core"
	"github.com/goboilerplates/micro-stream/api"
)

// SetAPIPaths sets API paths.
func SetAPIPaths(router *gin.Engine) {
	var getSamplesAPI api.GetSamplesAPI

	// inject GetSamplesAPIImpl .
	getSamplesAPI = api.GetSamplesAPIImpl{
		Interactor: core.CreateDefaultGetSamples("Kaka", "Ronaldo"),
	}

	apiV1 := router.Group("/api/v1")
	{
		apiV1.GET("/samples/all/:rateLimit", getSamplesAPI.All)
		apiV1.GET("/samples/name/:keyword/:rateLimit", getSamplesAPI.AllByName)
	}
}
