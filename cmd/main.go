package main

import (
	"health_check/internal/domain"
	"health_check/internal/servicies"
	"health_check/internal/servicies/check_service"
	"log"
)

func main() {
	inputService := servicies.NewJsonConfigFileInput("config.json")
	checkingService := check_service.CheckService{}
	outputService := servicies.StdOutput{}
	controller := domain.NewController(inputService, checkingService, outputService)
	err := controller.InitChecks()
	if err != nil {
		log.Fatal(err)
	}
}
