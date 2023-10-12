package main

import (
	"github.com/cesarcruzc/nearshore_test/internal"
	dc_nearshore "github.com/cesarcruzc/nearshore_test/internal/core/dc-nearshore"
	"github.com/cesarcruzc/nearshore_test/internal/core/device"
	"github.com/cesarcruzc/nearshore_test/internal/core/firmware"
	"github.com/cesarcruzc/nearshore_test/pkg/bootstrap"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	err := bootstrap.InitLoadEnv()
	if err != nil {
		log.Fatal(err)
	}

	logger := bootstrap.InitLogger()

	db, err := bootstrap.DBConnection()
	if err != nil {
		log.Fatal(err)
	}

	nearshoreEndpoints := dc_nearshore.MakeEndpoints()

	deviceRepository := device.NewRepository(logger, db)
	deviceService := device.NewService(logger, deviceRepository)
	deviceEndpoints := device.MakeEndpoints(deviceService)

	firmwareRepository := firmware.NewRepository(logger, db)
	firmwareService := firmware.NewService(logger, firmwareRepository)
	firmwareEndpoints := firmware.MakeEndpoints(firmwareService)

	gin.SetMode(gin.ReleaseMode)

	server := internal.LoadUrls(deviceEndpoints, firmwareEndpoints, nearshoreEndpoints)

	logger.Println("Server running on", ":8080")
	logger.Fatalf("Error in server: %s", server.Run(":8080"))
}
