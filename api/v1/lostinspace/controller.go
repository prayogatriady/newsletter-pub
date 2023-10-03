package lost

import (
	"fmt"
	"net/http"
	"newsletter-pub/http/httpresponse"

	"github.com/gin-gonic/gin"
)

func LostInSpace(c *gin.Context) {

	var errorMessage string

	lang := c.GetHeader("lang")
	if lang == "EN" {
		errorMessage = "You are lost in space"
	} else if lang == "ID" {
		errorMessage = "Anda tersesat"
	}

	httpresponse.BaseResponse(&httpresponse.HttpParams{
		GinContext:   c,
		StatusCode:   http.StatusOK,
		ServiceName:  "LostInSpace",
		ErrorMessage: fmt.Sprintf("No Route: %s", errorMessage),
	})
}
