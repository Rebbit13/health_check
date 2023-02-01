package main

import (
	"health_check/internal/domain"
	"health_check/internal/interfaces/alerts"
	repository2 "health_check/internal/interfaces/repository"
	"health_check/internal/servicies"
	"health_check/internal/servicies/check_service"
	"health_check/internal/servicies/output_service"
	"log"
)

func configureOutput() *output_service.MultiOutput {
	stdOutService := output_service.StdOutput{}
	entities := []interface{}{&repository2.SiteFullCheck{}}
	db := repository2.NewSqliteDatabase(entities)
	repository := repository2.NewGormRepository(db)
	sender := alerts.MockSender{}
	repositoryOutService := output_service.NewRepositoryOutput(repository, sender)
	return output_service.NewComplexOutput([]domain.OutputService{stdOutService, repositoryOutService})
}

func main() {
	inputService := servicies.NewJsonConfigFileInput("config.json")
	checkingService := check_service.CheckService{}
	outputService := configureOutput()
	controller := domain.NewController(inputService, checkingService, outputService)
	err := controller.InitChecks()
	if err != nil {
		log.Fatal(err)
	}
}
