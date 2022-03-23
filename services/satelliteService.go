package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/oskarincon/operation-quasar-go/constants"
	"github.com/oskarincon/operation-quasar-go/models"
	"github.com/patrickmn/go-cache"
)

var Cache *cache.Cache

func Init() {
	fmt.Printf("[PostInfoSatellite] - INIT")
	if Cache == nil {
		Cache = cache.New(50*time.Minute, 50*time.Minute)
	}
}

func PostInfoSatellite(satelliteName string, data models.Satellite) (response models.Satellite, err error) {
	data.Name = satelliteName
	fmt.Printf("[PostInfoSatellite] - data: %#v\n", data)
	SetCache(satelliteName, data)
	response = data
	return response, err
}

func SetCache(key string, emp models.Satellite) bool {
	Cache.Set(key, emp, cache.NoExpiration)
	return true
}

func GetCache(key string) (models.Satellite, bool) {
	var emp models.Satellite
	var found bool
	data, found := Cache.Get(key)
	if found {
		emp = data.(models.Satellite)
	}
	return emp, found
}

func GetInfoSatellite() (response models.Satellites, err error) {
	dataKenobi, foundKenobi := GetCache("kenobi")
	dataSkywalker, foundSkywalker := GetCache("skywalker")
	dataSato, foundSato := GetCache("sato")
	fmt.Printf("[GetInfoSatellite] - dataKenobi: %#v\n", dataKenobi)
	fmt.Printf("[GetInfoSatellite] - dataSkywalker: %#v\n", dataSkywalker)
	fmt.Printf("[GetInfoSatellite] - dataSato: %#v\n", dataSato)
	if !foundKenobi || !foundSkywalker || !foundSato {
		return response, errors.New(constants.DATA_NOT_FOUND_ERROR)
	}
	data := []models.Satellite{dataKenobi, dataSkywalker, dataSato}
	fmt.Printf("[GetInfoSatellite] - data res: %#v\n", data)
	response.Satellites = data
	return response, nil
}
