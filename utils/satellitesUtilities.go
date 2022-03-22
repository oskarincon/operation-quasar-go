package utils

import (
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/oskarincon/operation-quasar-go/constants"
	"github.com/oskarincon/operation-quasar-go/models"
)

func GetLocation(distances []float64) (position models.Position, err error) {
	_, errorValidate := validateDistances(distances)
	if errorValidate != nil {
		fmt.Printf("[GetLocation] errorValidate: %#v\n", errorValidate)
		return position, errorValidate
	}
	orb1 := models.Coordinate{X: constants.KENOBIX, Y: constants.KENOBIY, Distance: distances[0]}
	orb2 := models.Coordinate{X: constants.SKYWALKERX, Y: constants.SKYWALKERY, Distance: distances[1]}
	orb3 := models.Coordinate{X: constants.SATOX, Y: constants.SATOY, Distance: distances[2]}
	position = getDistance(orb1, orb2, orb3)
	return position, nil
}

func GetMessage(messages [][]string) (message string, err error) {
	errorValidate := validateMessages(messages)
	if errorValidate != nil {
		fmt.Printf("[GetMessage] errorValidate: %#v\n", errorValidate)
		return "", errorValidate
	}
	var message1, message2, message3 []string
	message1 = messages[0]
	message2 = messages[1]
	message3 = messages[2]
	if len(message1) != len(message2) || len(message1) != len(message3) || len(message2) != len(message3) {
		messagesFormated := deleteDesfase(message1, message2, message3)
		message1 = messagesFormated[0]
		message2 = messagesFormated[1]
		message3 = messagesFormated[2]
	}
	var mensajeFull string
	fmt.Printf("[GetMessage] len(message1): %#v\n", len(message1))
	for i := 0; i < len(message1); i++ {
		if !(message1[i] == "") {
			fmt.Printf("[GetMessage] message1[i]: %#v\n", message1[i])
			mensajeFull = validateWord(message1[i], mensajeFull)
		} else if !(message2[i] == "") {
			fmt.Printf("[GetMessage] message2[i]: %#v\n", message2[i])
			mensajeFull = validateWord(message2[i], mensajeFull)
		} else if !(message3[i] == "") {
			fmt.Printf("[GetMessage] message3[i]: %#v\n", message3[i])
			mensajeFull = validateWord(message3[i], mensajeFull)
		} else {
			if mensajeFull == "" {
				fmt.Printf("[GetMessage] error!!: %#v\n", mensajeFull)
				return "", errors.New(constants.MESSAGE_ERROR)
			}
		}
	}
	fmt.Printf("[GetMessage] mensajeFull: %#v\n", mensajeFull)
	return mensajeFull, nil
}

func getDistance(orb1, orb2, orb3 models.Coordinate) (position models.Position) {
	a := 2*orb2.X - 2*orb1.X
	b := 2*orb2.Y - 2*orb1.Y
	c := math.Pow(orb1.Distance, 2) - math.Pow(orb2.Distance, 2) - math.Pow(orb1.X, 2) + math.Pow(orb2.X, 2) - math.Pow(orb1.Y, 2) + math.Pow(orb2.Y, 2)
	d := 2*orb3.X - 2*orb2.X
	e := 2*orb3.Y - 2*orb2.Y
	f := math.Pow(orb2.Distance, 2) - math.Pow(orb3.Distance, 2) - math.Pow(orb2.X, 2) + math.Pow(orb3.X, 2) - math.Pow(orb2.Y, 2) + math.Pow(orb3.Y, 2)
	position.XCoordinate = (c*e - f*b) / (e*a - b*d)
	position.YCoordinate = (c*d - a*f) / (b*d - a*e)
	return position
}

func validateDistances(distances []float64) (bool, error) {
	for _, distance := range distances {
		if distance <= 0 {
			fmt.Printf("[GetLocation] validateDistances:  error distance < 0: %#v\n", distance)
			return false, errors.New(constants.DISTANCE_ERROR)
		}
	}
	return true, nil
}

func validateMessages(messages [][]string) error {
	if isEmptyMessage(messages[0]) || isEmptyMessage(messages[1]) || isEmptyMessage(messages[2]) {
		fmt.Printf("[GetMessage] validateMessages: error messages void: %#v\n", messages)
		return errors.New(constants.MESSAGE_ERROR)
	}
	return nil
}

func isEmptyMessage(arrays []string) bool {
	return arrays == nil || len(arrays) == 0
}

func deleteDesfase(message1 []string, message2 []string, message3 []string) [][]string {
	if len(message1) == len(message2) {
		if len(message3) < len(message2) {
			message2 = removeElementUsingCollection(message2, 1)
			message1 = removeElementUsingCollection(message1, 1)
		} else {
			message3 = removeElementUsingCollection(message3, 1)
		}
	} else if len(message2) == len(message3) {
		if len(message1) < len(message2) {
			message2 = removeElementUsingCollection(message2, 1)
			message3 = removeElementUsingCollection(message3, 1)
		} else {
			message2 = removeElementUsingCollection(message2, 1)
		}
	} else if len(message1) == len(message3) {
		if len(message2) < len(message3) {
			message3 = removeElementUsingCollection(message3, 1)
			message1 = removeElementUsingCollection(message1, 1)
		} else {
			message2 = removeElementUsingCollection(message2, 1)
		}
	}
	messages := make([][]string, 3, len(message1))
	messages[0] = message1
	messages[1] = message2
	messages[2] = message3
	return messages
}

func removeElementUsingCollection(arr []string, index int) []string {
	return append(arr[:index], arr[index+1:]...)
}

func validateWord(word string, message string) string {
	msg := strings.Split(message, "\\s+")
	for i := 0; i < len(msg); i++ {
		if word == msg[i] {
			return message + " "
		}
	}
	return message + word + " "
}
