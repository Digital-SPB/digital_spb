package data

import (
	"context"
	"encoding/json"
	"os"

	"github.com/greenblat17/digital_spb/internal/entity"
	"github.com/greenblat17/digital_spb/internal/service"
	"github.com/greenblat17/digital_spb/pkg/reader/csv"
	log "github.com/sirupsen/logrus"
)

func ScanEducationalDirection(services *service.Services) {
	configFile, err := os.Open("./assets/receiving_company.json")
	if err != nil {
		log.Info("opening config file", err.Error())
	}

	jsonParser := json.NewDecoder(configFile)
	var s []entity.EducatitionalDirection
	if err = jsonParser.Decode(&s); err != nil {
		log.Info("parsing config file", err.Error())
	}

	count, err := services.EducationalDirection.CountEducationalDirection(context.Background())
	if err != nil {
		log.Fatal("error scanning count educational direction")
	}
	log.Info("count: ", count)
	if count == 0 {
		log.Info("Initializing data...")
		for _, v := range s {
			services.EducationalDirection.CreateEducationalDirection(context.Background(), v)
		}
	}

}

func ScanVacancy(services *service.Services) {
	data, err := csv.ReadCsvFile("/assets/vacancy.csv")
	if err != nil {
		log.Error("error reading csv file: ", err.Error())
		return
	}

	count, err := services.Vacancy.CountVacancy(context.Background())
	if err != nil {
		log.Fatal("error scanning count vacancy")
	}

	if count == 0 {
		log.Info("Initializing data vacancy...")

		existVacancy := make(map[string]bool)
		for _, v := range data {
			if existVacancy[v[1]] {
				continue
			}
			vacancy := entity.Vacancy{
				Name: v[1],
			}
			existVacancy[v[1]] = true
			services.Vacancy.CreateVacancy(context.Background(), vacancy)
		}
	}

}
