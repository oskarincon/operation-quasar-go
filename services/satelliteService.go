package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/oskarincon/operation-quasar-go/cache"
	"github.com/oskarincon/operation-quasar-go/constants"
	"github.com/oskarincon/operation-quasar-go/models"
)

func PostInfoSatellite(satelliteName string, data models.Satellite) (response models.Satellite, err error) {
	data.Name = satelliteName
	fmt.Printf("[PostInfoSatellite] - data: %#v\n", data)
	cache.DefaultMemCache.Set(satelliteName, data, 60*time.Minute)
	response = data
	return response, err
}

func GetInfoSatellite() (response models.Satellites, err error) {
	dataKenobi, foundKenobi := cache.DefaultMemCache.Get("kenobi")
	dataSkywalker, foundSkywalker := cache.DefaultMemCache.Get("skywalker")
	dataSato, foundSato := cache.DefaultMemCache.Get("sato")
	if !foundKenobi || !foundSkywalker || !foundSato {
		return response, errors.New(constants.DATA_NOT_FOUND_ERROR)
	}
	data := []models.Satellite{dataKenobi.(models.Satellite), dataSkywalker.(models.Satellite), dataSato.(models.Satellite)}
	fmt.Printf("[GetInfoSatellite] - data res: %#v\n", data)
	response.Satellites = data
	return response, nil
}
