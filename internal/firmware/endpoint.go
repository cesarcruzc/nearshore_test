package firmware

import (
	"fmt"
	"github.com/cesarcruzc/nearshore_test/pkg/meta"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type (
	Controller func(ctx *gin.Context)

	Endpoints struct {
		Create Controller
		Get    Controller
		GetAll Controller
		Update Controller
		Delete Controller
	}

	CreateRequest struct {
		Name         string `json:"name"`
		DeviceId     string `json:"device_id"`
		Version      string `json:"version"`
		ReleaseNotes string `json:"release_notes"`
		ReleaseDate  string `json:"release_date"`
		Url          string `json:"url"`
	}

	UpdateRequest struct {
		Name         *string `json:"name"`
		DeviceId     *string `json:"device_id"`
		Version      *string `json:"version"`
		ReleaseNotes *string `json:"release_notes"`
		ReleaseDate  *string `json:"release_date"`
		Url          *string `json:"url"`
	}

	Response struct {
		Status int         `json:"status"`
		Data   interface{} `json:"data,omitempty"`
		Err    string      `json:"error,omitempty"`
		Meta   meta.Meta   `json:"meta,omitempty"`
	}
)

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		Create: makeCreateEndpoint(s),
		Get:    makeGetEndpoint(s),
		GetAll: makeGetAllEndpoint(s),
		Update: makeUpdateEndpoint(s),
		Delete: makeDeleteEndpoint(s),
	}
}

func makeCreateEndpoint(s Service) Controller {
	return func(ctx *gin.Context) {
		var req CreateRequest

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if req.Name == "" {
			ctx.JSON(http.StatusBadRequest, &Response{Status: http.StatusBadRequest, Err: fmt.Sprintf("name is required")})
			return
		}

		if req.DeviceId == "" {
			ctx.JSON(http.StatusBadRequest, &Response{Status: http.StatusBadRequest, Err: fmt.Sprintf("device_id is required")})
			return
		}

		if req.Version == "" {
			ctx.JSON(http.StatusBadRequest, &Response{Status: http.StatusBadRequest, Err: fmt.Sprintf("version is required")})
			return
		}

		if req.ReleaseNotes == "" {
			ctx.JSON(http.StatusBadRequest, &Response{Status: http.StatusBadRequest, Err: fmt.Sprintf("release_notes is required")})
			return
		}

		if req.ReleaseDate == "" {
			ctx.JSON(http.StatusBadRequest, &Response{Status: http.StatusBadRequest, Err: fmt.Sprintf("release_date is required")})
			return
		}

		if req.Url == "" {
			ctx.JSON(http.StatusBadRequest, &Response{Status: http.StatusBadRequest, Err: fmt.Sprintf("url is required")})
			return
		}

		firmware, err := s.Create(req.Name, req.DeviceId, req.Version, req.ReleaseNotes, req.ReleaseDate, req.Url)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &Response{Status: http.StatusInternalServerError, Err: err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, &Response{Status: http.StatusCreated, Data: firmware})

	}
}

func makeGetAllEndpoint(s Service) Controller {
	return func(ctx *gin.Context) {
		//v := r.URL.Query()

		v := ctx.Request.URL.Query()

		filters := Filters{
			Name: v.Get("name"),
		}

		limit, err := strconv.ParseInt(v.Get("limit"), 10, 64)
		if err != nil {
			limit = 0
		}

		page, err := strconv.ParseInt(v.Get("page"), 10, 64)
		if err != nil {
			page = 0
		}

		count, err := s.Count(filters)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &Response{Status: http.StatusInternalServerError, Err: err.Error()})
			return
		}

		metaData, err := meta.New(page, limit, count)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &Response{Status: http.StatusInternalServerError, Err: err.Error()})
			return
		}

		firmwares, err := s.GetAll(filters, metaData.Offset(), metaData.Limit())
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &Response{Status: http.StatusInternalServerError, Err: err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, &Response{Status: http.StatusOK, Data: firmwares, Meta: *metaData})
	}
}

func makeGetEndpoint(s Service) Controller {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		firmware, err := s.Get(id)

		if err != nil {
			ctx.JSON(http.StatusNotFound, &Response{Status: http.StatusNotFound, Err: err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, &Response{Status: http.StatusOK, Data: firmware})
	}
}

func makeUpdateEndpoint(s Service) Controller {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		var req UpdateRequest

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, &Response{Status: http.StatusBadRequest, Err: "Invalid request format"})
			return
		}

		if req.Name != nil && *req.Name == "" {
			ctx.JSON(http.StatusBadRequest, &Response{Status: http.StatusBadRequest, Err: fmt.Sprintf("Name cannot be empty")})
			return
		}

		if req.DeviceId != nil && *req.DeviceId == "" {
			ctx.JSON(http.StatusBadRequest, &Response{Status: http.StatusBadRequest, Err: fmt.Sprintf("DeviceId cannot be empty")})
			return
		}

		if req.Version != nil && *req.Version == "" {
			ctx.JSON(http.StatusBadRequest, &Response{Status: http.StatusBadRequest, Err: fmt.Sprintf("Version cannot be empty")})
			return
		}

		if req.ReleaseNotes != nil && *req.ReleaseNotes == "" {
			ctx.JSON(http.StatusBadRequest, &Response{Status: http.StatusBadRequest, Err: fmt.Sprintf("ReleaseNotes cannot be empty")})
			return
		}

		if req.ReleaseDate != nil && *req.ReleaseDate == "" {
			ctx.JSON(http.StatusBadRequest, &Response{Status: http.StatusBadRequest, Err: fmt.Sprintf("ReleaseDate cannot be empty")})
			return
		}

		if req.Url != nil && *req.Url == "" {
			ctx.JSON(http.StatusBadRequest, &Response{Status: http.StatusBadRequest, Err: fmt.Sprintf("Url cannot be empty")})
			return
		}

		err := s.Update(id, req.Name, req.DeviceId, req.Version, req.ReleaseNotes, req.ReleaseDate, req.Url)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &Response{Status: http.StatusInternalServerError, Err: err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, &Response{Status: http.StatusOK, Data: "success"})
	}
}

func makeDeleteEndpoint(s Service) Controller {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		err := s.Delete(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &Response{Status: http.StatusInternalServerError, Err: err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, &Response{Status: http.StatusOK, Data: "success"})
	}
}
