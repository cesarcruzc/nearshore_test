package internal

import (
	dc_nearshore "github.com/cesarcruzc/nearshore_test/internal/core/dc-nearshore"
	"github.com/cesarcruzc/nearshore_test/internal/core/device"
	"github.com/cesarcruzc/nearshore_test/internal/core/firmware"
	"github.com/gin-gonic/gin"
)

func LoadUrls(username, password string, deviceEndpoints device.Endpoints, firmwareEndpoints firmware.Endpoints, nearshoreEndpoints dc_nearshore.Endpoints) *gin.Engine {
	router := gin.New()

	router.GET("/", gin.HandlerFunc(nearshoreEndpoints.Root))
	router.GET("/health", gin.HandlerFunc(nearshoreEndpoints.HealthCheck))

	router.Use(BasicAuthMiddleware(username, password))
	deviceGroup := router.Group("/device")
	{
		deviceGroup.POST("/", gin.HandlerFunc(deviceEndpoints.Create))
		deviceGroup.GET("/", gin.HandlerFunc(deviceEndpoints.GetAll))
		deviceGroup.GET("/:id", gin.HandlerFunc(deviceEndpoints.Get))
		deviceGroup.PUT("/:id", gin.HandlerFunc(deviceEndpoints.Update))
		deviceGroup.DELETE("/:id", gin.HandlerFunc(deviceEndpoints.Delete))
	}

	firmwareGroup := router.Group("/firmware")
	{
		firmwareGroup.POST("/", gin.HandlerFunc(firmwareEndpoints.Create))
		firmwareGroup.GET("/", gin.HandlerFunc(firmwareEndpoints.GetAll))
		firmwareGroup.GET("/:id", gin.HandlerFunc(firmwareEndpoints.Get))
		firmwareGroup.PUT("/:id", gin.HandlerFunc(firmwareEndpoints.Update))
		firmwareGroup.DELETE("/:id", gin.HandlerFunc(firmwareEndpoints.Delete))
	}

	return router
}
