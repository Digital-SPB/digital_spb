package data

import (
	"encoding/json"
	"os"

	"github.com/greenblat17/digital_spb/internal/entity"
	log "github.com/sirupsen/logrus"
)

func ScanEducationalDirection() []entity.EducatitionalDirection {
	configFile, err := os.Open("./assets/receiving_company.json")
	if err != nil {
		log.Info("opening config file", err.Error())
	}

	jsonParser := json.NewDecoder(configFile)
	var s []entity.EducatitionalDirection
	if err = jsonParser.Decode(&s); err != nil {
		log.Info("parsing config file", err.Error())
	}

	log.Printf("scan complete %s", s[0].Name)

	return s
}
