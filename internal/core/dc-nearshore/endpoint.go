package dc_nearshore

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	Controller func(ctx *gin.Context)

	Endpoints struct {
		Root        Controller
		HealthCheck Controller
	}

	Response struct {
		Status int         `json:"status"`
		Data   interface{} `json:"data,omitempty"`
		Err    string      `json:"error,omitempty"`
	}
)

func MakeEndpoints() Endpoints {
	return Endpoints{
		Root:        makeRootEndpoint(),
		HealthCheck: makeHealthCheckEndpoint(),
	}
}

func makeRootEndpoint() Controller {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, &Response{
			Status: http.StatusOK,
			Data:   "Nearshore Test",
		})
	}
}

func makeHealthCheckEndpoint() Controller {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, &Response{
			Status: http.StatusOK,
			Data:   "ok",
		})
	}
}
