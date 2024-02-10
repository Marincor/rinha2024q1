package logging

import (
	"fmt"

	"api.default.marincor.com/entity"
)

func Log(details *entity.LogDetails, severity string, resourceLabels *map[string]string) {
	fmt.Sprintf("%s", details.Message)
}
