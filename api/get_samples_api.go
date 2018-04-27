package api

import (
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/goboilerplates/core"
	"github.com/gonitor/gonitor/util"
)

// GetSamplesAPI is the interface for GetSamplesAPI.
type GetSamplesAPI interface {
	All(context *gin.Context)
	AllByName(context *gin.Context)
}

// GetSamplesAPIImpl is the implementation of GetSamplesAPI interface.
type GetSamplesAPIImpl struct {
	Interactor core.GetSamplesInteractor
}

// All gets all samples.
func (api GetSamplesAPIImpl) All(context *gin.Context) {
	rateLimit, convertErr := util.ConvertStringToTimeDuration(context.Param("rateLimit"))
	context.Stream(func(w io.Writer) bool {
		if convertErr != nil {
			return false
		}
		time.Sleep(rateLimit * time.Second)
		body, err := api.Interactor.All()
		return HandleStreamResponse(context, body, err, "GetSamplesAll")
	})
}

// AllByName gets all samples that have name matched with keyword.
func (api GetSamplesAPIImpl) AllByName(context *gin.Context) {
	keyword := context.Param("keyword")
	rateLimit, convertErr := util.ConvertStringToTimeDuration(context.Param("rateLimit"))
	context.Stream(func(w io.Writer) bool {
		if convertErr != nil {
			return false
		}
		time.Sleep(rateLimit * time.Second)
		body, err := api.Interactor.AllByName(keyword)
		return HandleStreamResponse(context, body, err, "GetSamplesAllByName")
	})
}
