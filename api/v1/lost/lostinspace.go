package lost

import (
	"net/http"
	"newsletter-pub/http/httpresponse"

	"github.com/gin-gonic/gin"
)

func LostInSpace(c *gin.Context) {

	httpresponse.BaseResponse(&httpresponse.HttpParams{
		GinContext:   c,
		StatusCode:   http.StatusOK,
		ServiceName:  "LostInSpace",
		ErrorMessage: "You are lost in space",
	})
}
