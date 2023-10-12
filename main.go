package main

import (
	"github.com/cesarcruzc/nearshore_test/internal/device"
	"github.com/cesarcruzc/nearshore_test/internal/firmware"
	"github.com/cesarcruzc/nearshore_test/pkg/bootstrap"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()

	err := bootstrap.InitLoadEnv()
	if err != nil {
		log.Fatal(err)
	}

	logger := bootstrap.InitLogger()

	db, err := bootstrap.DBConnection()
	if err != nil {
		log.Fatal(err)
	}

	deviceRepository := device.NewRepository(logger, db)
	deviceService := device.NewService(logger, deviceRepository)
	deviceEndpoints := device.MakeEndpoints(deviceService)

	firmwareRepository := firmware.NewRepository(logger, db)
	firmwareService := firmware.NewService(logger, firmwareRepository)
	firmwareEndpoints := firmware.MakeEndpoints(firmwareService)

	router.POST("/device", gin.HandlerFunc(deviceEndpoints.Create))
	router.GET("/device", gin.HandlerFunc(deviceEndpoints.GetAll))
	router.GET("/device/:id", gin.HandlerFunc(deviceEndpoints.Get))
	router.PUT("/device/:id", gin.HandlerFunc(deviceEndpoints.Update))
	router.DELETE("/device/:id", gin.HandlerFunc(deviceEndpoints.Delete))

	router.POST("/firmware", gin.HandlerFunc(firmwareEndpoints.Create))
	router.GET("/firmware", gin.HandlerFunc(firmwareEndpoints.GetAll))
	router.GET("/firmware/:id", gin.HandlerFunc(firmwareEndpoints.Get))
	router.PUT("/firmware/:id", gin.HandlerFunc(firmwareEndpoints.Update))
	router.DELETE("/firmware/:id", gin.HandlerFunc(firmwareEndpoints.Delete))

	router.Run(":8080")

}
