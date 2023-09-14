package httpresponse

import (
	"encoding/json"
	"log"
	"net/http"
	"newsletter-pub/utils/config"
	"time"

	"github.com/gin-gonic/gin"
)

type HttpParams struct {
	GinContext   *gin.Context
	Data         interface{}
	Payload      interface{}
	StatusCode   int
	ServiceName  string
	ErrorMessage string
}

type Response struct {
	Status       int         `json:"status"`
	Data         interface{} `json:"data"`
	ErrorMessage string      `json:"error_message"`
	AppName      string      `json:"app_name"`
	AppVersion   string      `json:"app_version"`
	CurrentTime  string      `json:"current_time"`
}

var appConfig = config.AppCfg

func BaseResponse(httpParams *HttpParams) {

	loc, _ := time.LoadLocation(appConfig.App.Timezone)
	currentTime := time.Now().In(loc)
	currentTimeNew := currentTime.Format(time.RFC3339)

	payload, err := json.Marshal(httpParams.Payload)
	if err != nil {
		return
	}
	data, err := json.Marshal(httpParams.Data)
	if err != nil {
		return
	}

	log.Printf("%s | %s | %s | %d | %s | %s | %s \n",
		currentTimeNew,
		httpParams.ServiceName,
		httpParams.GinContext.Request.Host+httpParams.GinContext.Request.URL.Path,
		httpParams.StatusCode,
		payload,
		data,
		httpParams.ErrorMessage)

	response := Response{
		Status:       httpParams.StatusCode,
		Data:         httpParams.Data,
		ErrorMessage: httpParams.ErrorMessage,
		AppName:      appConfig.App.AppName,
		AppVersion:   appConfig.App.AppVersion,
		CurrentTime:  currentTimeNew,
	}

	httpParams.GinContext.JSON(http.StatusOK, &response)
}
