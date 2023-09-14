package email

import (
	"fmt"
	"net/http"
	"newsletter-pub/http/httpresponse"
	models_email "newsletter-pub/models/email"

	"github.com/gin-gonic/gin"
)

type EmailController interface {
	SendEmail(c *gin.Context)
}

type emailController struct {
	emailService EmailService
}

// Dependency injection
func NewEmailController(service EmailService) EmailController {
	return &emailController{
		emailService: service,
	}
}

func (ctrl *emailController) SendEmail(c *gin.Context) {

	serviceName := "SendEmail"

	var payload models_email.EmailPayload
	if err := c.BindJSON(&payload); err != nil {
		httpresponse.BaseResponse(&httpresponse.HttpParams{
			GinContext:   c,
			Payload:      payload,
			StatusCode:   http.StatusBadRequest,
			ServiceName:  serviceName,
			ErrorMessage: fmt.Sprintf("Bad Request: %s", err),
		})
		return
	}

	response, err := ctrl.emailService.SendEmail(&payload)
	if err != nil {
		httpresponse.BaseResponse(&httpresponse.HttpParams{
			GinContext:   c,
			Payload:      payload,
			StatusCode:   http.StatusInternalServerError,
			ServiceName:  serviceName,
			ErrorMessage: fmt.Sprintf("Something went wrong: %s", err),
		})
		return
	}

	httpresponse.BaseResponse(&httpresponse.HttpParams{
		GinContext:  c,
		Data:        response,
		Payload:     payload,
		StatusCode:  http.StatusOK,
		ServiceName: serviceName,
	})
}
