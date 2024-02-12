package logging

import (
	"log"

	"api.default.marincor.com/entity"
)

func Log(details *entity.LogDetails, severity string, resourceLabels *map[string]string) {
	log.Println(details)
	log.Println(severity)
	log.Println(resourceLabels)
}
