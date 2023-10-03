package ping

import (
	"fmt"
	"net/http"
	"newsletter-pub/http/httpresponse"

	"github.com/gin-gonic/gin"
)

type PingController interface {
	Ping(c *gin.Context)
}

type pingController struct {
	pingService PingService
}

// Dependency injection
func NewPingController(service PingService) PingController {
	return &pingController{
		pingService: service,
	}
}

func (ctrl *pingController) Ping(c *gin.Context) {

	serviceName := "Ping"

	response, err := ctrl.pingService.Ping()
	if err != nil {
		httpresponse.BaseResponse(&httpresponse.HttpParams{
			GinContext:   c,
			StatusCode:   http.StatusInternalServerError,
			ServiceName:  serviceName,
			ErrorMessage: fmt.Sprintf("Something went wrong: %s", err),
		})
		return
	}

	httpresponse.BaseResponse(&httpresponse.HttpParams{
		GinContext:  c,
		Data:        response,
		StatusCode:  http.StatusOK,
		ServiceName: serviceName,
	})
}
