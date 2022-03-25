package services

import (
	"errors"
	"fmt"

	"github.com/oskarincon/operation-quasar-go/constants"
	"github.com/oskarincon/operation-quasar-go/models"
	"github.com/oskarincon/operation-quasar-go/utils"
)

func FindPositionMessage(data models.Satellites) (response models.Response, err error) {
	var distanceKenobi, distanceSkywalker, distanceSato float64
	var messagesKenobi, messagesSkywalker, messagesSato []string
	for _, satellite := range data.Satellites {
		if satellite.Name == constants.KENOBI {
			distanceKenobi = satellite.Distance
			messagesKenobi = satellite.Message
		}
		if satellite.Name == constants.SKYWALKER {
			distanceSkywalker = satellite.Distance
			messagesSkywalker = satellite.Message
		}
		if satellite.Name == constants.SATO {
			distanceSato = satellite.Distance
			messagesSato = satellite.Message
		}
	}
	if len(messagesKenobi) <= 0 || len(messagesSkywalker) <= 0 || len(messagesSato) <= 0 {
		return response, errors.New(constants.DISTANCES_SATELLITES_ERROR)
	}
	distances := []float64{distanceKenobi, distanceSkywalker, distanceSato}
	messages := [][]string{messagesKenobi, messagesSkywalker, messagesSato}
	fmt.Printf("[FindPositionMessage] distances: %#v\n", distances)
	pos, err := utils.GetLocation(distances)
	fmt.Printf("[FindPositionMessage] messages: %#v\n", messages)
	msg, err := utils.GetMessage(messages)
	response.Message = msg
	response.Position = pos
	fmt.Printf("[FindPositionMessage] response: %#v\n, err: %#v\n", response, err)
	return response, err
}
