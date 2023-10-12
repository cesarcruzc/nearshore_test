package main

import (
	"github.com/cesarcruzc/nearshore_test/internal/device"
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

	router.POST("/device", gin.HandlerFunc(deviceEndpoints.Create))
	router.GET("/device", gin.HandlerFunc(deviceEndpoints.GetAll))
	router.GET("/device/:id", gin.HandlerFunc(deviceEndpoints.Get))
	router.PUT("/device/:id", gin.HandlerFunc(deviceEndpoints.Update))
	router.DELETE("/device/:id", gin.HandlerFunc(deviceEndpoints.Delete))

	router.Run(":8080")

}
