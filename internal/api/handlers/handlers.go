package handlers

import (
	"golang-api/internal/api/constants"
	"golang-api/pkg/ahttp"
	"golang-api/pkg/alog"
	"os"

	"github.com/kataras/iris/v12"
)

func httpSuccess(c iris.Context, data interface{}, message string) {
	response := ahttp.Response{}
	response.Status = ahttp.ResponseSuccess
	response.Message = message

	if data != nil {
		response.Data = data
	}

	c.JSON(response)
	c.StopExecution()
}

func httpError(c iris.Context, httpError ahttp.ErrorResponse, err error) {
	c.StatusCode(httpError.Status)

	alog.Logger.Error(err.Error())

	if os.Getenv(constants.AppEnv) == constants.EnvDevelopment {
		traces := alog.GetTracer(err)

		response := ahttp.ErrorDebugResponse{
			Status:  httpError.Status,
			Code:    httpError.Code,
			Message: err.Error(),
			Debug:   traces,
		}

		c.JSON(response)
	} else {
		response := ahttp.ErrorResponse{
			Status:  httpError.Status,
			Code:    httpError.Code,
			Message: err.Error(),
		}

		c.JSON(response)
	}

	c.StopExecution()
}
